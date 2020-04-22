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

		// String "testing"
		{
			data: []byte("\x12\x07\x74\x65\x73\x74\x69\x6e\x67"),
			expected: Message(map[string]interface{}{
				"__2": "testing",
			}),
			err: nil,
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
