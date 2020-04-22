package protosl

import (
	"errors"
	"fmt"
	"google.golang.org/protobuf/encoding/protowire"
	"unicode/utf8"
)

var (
	errReserved   = errors.New("cannot parse reserved wire type")
	errDeprecated = errors.New("deprecated wire type")
)

type Message map[string]interface{}

func getTag(num protowire.Number) string {
	return fmt.Sprintf("__%d", num)
}

func Unmarshal(b []byte, msg Message) error {
	for len(b) > 0 {
		// Parse the tag (field number and wire type).
		num, wtyp, tagLen := protowire.ConsumeTag(b)
		if tagLen < 0 {
			return protowire.ParseError(tagLen)
		}
		b = b[tagLen:]

		k := getTag(num)

		switch wtyp {
		case protowire.VarintType:
			v, n := protowire.ConsumeVarint(b)
			if err := protowire.ParseError(n); err != nil {
				return err
			}
			msg[k] = v
			b = b[n:]

		case protowire.Fixed32Type:
			v, n := protowire.ConsumeFixed32(b)
			if err := protowire.ParseError(n); err != nil {
				return err
			}
			msg[k] = v
			b = b[n:]

		case protowire.Fixed64Type:
			v, n := protowire.ConsumeFixed64(b)
			if err := protowire.ParseError(n); err != nil {
				return err
			}
			msg[k] = v
			b = b[n:]

		case protowire.BytesType:
			d, n := protowire.ConsumeBytes(b)
			if err := protowire.ParseError(n); err != nil {
				return err
			}
			subMsg := Message{}

			if utf8.Valid(d) { // string
				msg[k] = string(d)
			} else if err := Unmarshal(d, subMsg); err == nil { // embedded message
				msg[k] = subMsg
				// else if ... TODO repeated
			} else { // bytes
				msg[k] = d
			}
			b = b[n:]

		case protowire.StartGroupType:
			// https://developers.google.com/protocol-buffers/docs/encoding#structure
			return errDeprecated

		case protowire.EndGroupType:
			// https://developers.google.com/protocol-buffers/docs/encoding#structure
			return errDeprecated

		default:
			return errReserved
		}
	}

	return nil
}
