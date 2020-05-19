package protodump

import (
	"errors"
	"fmt"
	"unicode/utf8"

	"google.golang.org/protobuf/encoding/protowire"
)

const (
	stringUnionKey  = "__string"
	bytesUnionKey   = "__bytes"
	packedUnionKey  = "__packed"
	messageUnionKey = "__message"
)

var (
	errReserved   = errors.New("cannot parse reserved wire type")
	errDeprecated = errors.New("deprecated wire type")
)

type Message map[int32]interface{}

type bytesUnion struct {
	bytes []byte

	string  string
	packed  []interface{}
	message Message
}

func newBytesUnion(b []byte) *bytesUnion {
	return &bytesUnion{
		bytes: b,
	}
}
func (b *bytesUnion) value() interface{} {
	rv := map[string]interface{}{
		bytesUnionKey: b.bytes,
	}

	if b.string != "" {
		rv[stringUnionKey] = b.string
	}
	if b.packed != nil {
		rv[packedUnionKey] = b.packed
	}
	if b.message != nil {
		rv[messageUnionKey] = b.message
	}

	return rv
}

func Unmarshal(b []byte, msg Message) error {
	for len(b) > 0 {
		// Parse the tag (field number and wire type).
		num, wtyp, tagLen := protowire.ConsumeTag(b)
		if tagLen < 0 {
			return protowire.ParseError(tagLen)
		}
		b = b[tagLen:]

		var v interface{}
		n := -1

		switch wtyp {
		case protowire.VarintType:
			v, n = protowire.ConsumeVarint(b)

		case protowire.Fixed32Type:
			v, n = protowire.ConsumeFixed32(b)

		case protowire.Fixed64Type:
			v, n = protowire.ConsumeFixed64(b)

		case protowire.BytesType:
			d, dl := protowire.ConsumeBytes(b)
			if err := protowire.ParseError(dl); err != nil {
				return fmt.Errorf("wire type %v, field number %v caused : %w", wtyp, num, err)
			}

			union := newBytesUnion(d)

			subMsg := Message{}
			if err := Unmarshal(d, subMsg); err == nil { // embedded message
				union.message = subMsg
			}

			if packed, err := extractPacked(d); err == nil { // packed repeated
				union.packed = packed
			}

			if utf8.Valid(d) { // string
				union.string = string(d)
			}

			n = dl
			v = union.value()

		case protowire.StartGroupType:
			// https://developers.google.com/protocol-buffers/docs/encoding#structure
			return fmt.Errorf("wire type %v, field number %v caused : %w", wtyp, num, errDeprecated)

		case protowire.EndGroupType:
			// https://developers.google.com/protocol-buffers/docs/encoding#structure
			return fmt.Errorf("wire type %v, field number %v caused : %w", wtyp, num, errDeprecated)

		default:
			return fmt.Errorf("wire type %v, field number %v caused : %w", wtyp, num, errReserved)
		}

		if err := protowire.ParseError(n); err != nil {
			return fmt.Errorf("wire type %v, field number %v caused : %w", wtyp, num, err)
		}

		// If the value appears twice or uppper, wrap it by slice
		k := int32(num)
		if e, ok := msg[k]; ok {
			switch vv := e.(type) {
			case []interface{}:
				msg[k] = append(vv, v)
			default:
				msg[k] = []interface{}{vv, v}
			}
		} else {
			msg[k] = v
		}

		b = b[n:]
	}

	return nil
}

func extractPacked(b []byte) ([]interface{}, error) {
	rv := []interface{}{}

	for len(b) > 0 {
		var v interface{}
		n := -1

		v, n = protowire.ConsumeVarint(b)
		if err := protowire.ParseError(n); err == nil {
			rv = append(rv, v)
			b = b[n:]
			continue
		}

		v, n = protowire.ConsumeFixed32(b)
		if err := protowire.ParseError(n); err == nil {
			rv = append(rv, v)
			b = b[n:]
			continue
		}

		v, n = protowire.ConsumeFixed64(b)
		if err := protowire.ParseError(n); err == nil {
			rv = append(rv, v)
			b = b[n:]
			continue
		}

		return nil, fmt.Errorf("non packed repeated numeric")
	}

	return rv, nil
}
