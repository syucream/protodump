package protosl

import (
	"io"
	"reflect"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	cases := []struct {
		data     []byte
		expected Message
		err      error
	}{
		// Varint 150
		{
			data: []byte("\x08\x96\x01"),
			expected: Message(map[string]interface{}{
				"__1": uint64(150),
			}),
			err: nil,
		},

		// Fixed64
		{
			data: []byte("\x09\x87\x96\xa5\xb4\xc3\xd2\xe1\xf0"),
			expected: Message(map[string]interface{}{
				"__1": uint64(0xf0e1d2c3b4a59687),
			}),
			err: nil,
		},

		// String "testing"
		{
			data: []byte("\x12\x07\x74\x65\x73\x74\x69\x6e\x67"),
			expected: Message(map[string]interface{}{
				"__2": "testing",
			}),
			err: nil,
		},

		// Fixed32
		{
			data: []byte("\x15\xc3\xd2\xe1\xf0"),
			expected: Message(map[string]interface{}{
				"__2": uint32(0xf0e1d2c3),
			}),
			err: nil,
		},

		// Start group
		{
			data:     []byte("\x0b\x00"),
			expected: Message{},
			err:      errDeprecated,
		},

		// End group
		{
			data:     []byte("\x0c\x00"),
			expected: Message{},
			err:      errDeprecated,
		},

		// Embedded message {"__1": 150}
		{
			data: []byte("\x1a\x03\x08\x96\x01"),
			expected: Message(map[string]interface{}{
				"__3": Message(map[string]interface{}{
					"__1": uint64(150),
				}),
			}),
			err: nil,
		},

		// TODO fixed32

		// TODO fixed64

		// TODO repeated

		// no field value
		{
			data:     []byte("\xff"),
			expected: Message{},
			err:      io.ErrUnexpectedEOF,
		},
	}

	for _, c := range cases {
		actual := Message{}
		err := Unmarshal(c.data, actual)

		if err != c.err {
			t.Errorf("expected: %v, but actual: %v", c.err, err)
		}

		if !reflect.DeepEqual(actual, c.expected) {
			t.Errorf("expected: %v, but actual: %v", c.expected, actual)
		}
	}
}
