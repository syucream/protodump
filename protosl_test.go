package protosl

import (
	"io"
	"reflect"
	"testing"

	"github.com/golang/protobuf/proto"
	protosl "github.com/syucream/protosl/internal"
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
			expected: Message(map[int32]interface{}{
				1: uint64(150),
			}),
			err: nil,
		},

		// Fixed64
		{
			data: []byte("\x09\x87\x96\xa5\xb4\xc3\xd2\xe1\xf0"),
			expected: Message(map[int32]interface{}{
				1: uint64(0xf0e1d2c3b4a59687),
			}),
			err: nil,
		},

		// String "testing"
		{
			data: []byte("\x12\x07\x74\x65\x73\x74\x69\x6e\x67"),
			expected: Message(map[int32]interface{}{
				2: (&bytesUnion{
					bytes:  []byte("\x74\x65\x73\x74\x69\x6e\x67"),
					string: "testing",
					packed: []interface{}{uint64(0x74), uint64(0x65), uint64(0x73), uint64(0x74), uint64(0x69), uint64(0x6e), uint64(0x67)},
				}).value(),
			}),
			err: nil,
		},

		// Fixed32
		{
			data: []byte("\x15\xc3\xd2\xe1\xf0"),
			expected: Message(map[int32]interface{}{
				2: uint32(0xf0e1d2c3),
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
			expected: Message(map[int32]interface{}{
				3: (&bytesUnion{
					bytes: []byte("\x08\x96\x01"),
					message: Message(map[int32]interface{}{
						1: uint64(150),
					}),
					packed: []interface{}{uint64(0x8), uint64(0x96)},
				}).value(),
			}),
			err: nil,
		},

		// Complex
		{
			data: func() []byte {
				msg := &protosl.Example{
					Uint64Val:  1,
					StringVal:  "testing",
					Fixed64Val: 11,
					Fixed32Val: 111,
					EnumVal:    protosl.Example_ONE,
					ChildVal: &protosl.Child{
						V: 1,
					},
					RUint64Val: []uint64{2, 3},
					RStringVal: []string{"aaa", "bbb"},
					// RFixed64Val: []uint64{22, 33}, TODO repeated fixed isn't supported
					// RFixed32Val: []uint32{222, 333}, TODO repeated fixed isn't supported
					REnumVal: []protosl.Example_Num{protosl.Example_ZERO, protosl.Example_ONE},
					RChildVal: []*protosl.Child{
						{
							V: 2,
						},
						{
							V: 3,
						},
					},
				}

				d, err := proto.Marshal(msg)
				if err != nil {
					t.Fatal(err)
				}

				return d
			}(),
			expected: Message(map[int32]interface{}{
				1: uint64(1),
				2: (&bytesUnion{
					bytes:  []byte("\x74\x65\x73\x74\x69\x6e\x67"),
					string: "testing",
					packed: []interface{}{uint64(0x74), uint64(0x65), uint64(0x73), uint64(0x74), uint64(0x69), uint64(0x6e), uint64(0x67)},
				}).value(),
				3: uint64(11),
				4: uint32(111),
				5: uint64(1),
				6: (&bytesUnion{
					bytes:  []byte("\x08\x01"),
					string: "\x08\x01",
					message: Message(map[int32]interface{}{
						1: uint64(1),
					}),
					packed: []interface{}{uint64(0x8), uint64(0x1)},
				}).value(),
				101: (&bytesUnion{
					bytes:  []byte("\x02\x03"),
					string: "\x02\x03",
					packed: []interface{}{uint64(2), uint64(3)},
				}).value(),
				102: []interface{}{
					(&bytesUnion{
						bytes:  []byte("\x61\x61\x61"),
						string: "aaa",
						packed: []interface{}{uint64(0x61), uint64(0x61), uint64(0x61)},
					}).value(),
					(&bytesUnion{
						bytes:  []byte("\x62\x62\x62"),
						string: "bbb",
						packed: []interface{}{uint64(0x62), uint64(0x62), uint64(0x62)},
					}).value(),
				},
				// 103: TODO repeated fixed isn't supported
				// 104: TODO repeated fixed isn't supported
				105: (&bytesUnion{
					bytes:  []byte("\x00\x01"),
					string: "\x00\x01",
					packed: []interface{}{uint64(0), uint64(1)},
				}).value(),
				106: []interface{}{
					(&bytesUnion{
						bytes:  []byte("\x08\x02"),
						string: "\x08\x02",
						message: Message(map[int32]interface{}{
							1: uint64(2),
						}),
						packed: []interface{}{uint64(0x08), uint64(0x02)},
					}).value(),
					(&bytesUnion{
						bytes:  []byte("\x08\x03"),
						string: "\x08\x03",
						message: Message(map[int32]interface{}{
							1: uint64(3),
						}),
						packed: []interface{}{uint64(0x08), uint64(0x03)},
					}).value(),
				},
			}),
			err: nil,
		},

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
