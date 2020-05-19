# protodump

A protobuf schema-less module.

## Summary

[Protocol Buffer Wire Format](https://developers.google.com/protocol-buffers/docs/encoding) is super simple and contains wire type in each field. Generally it expects descriptors on marshaler/unmarshaler but this module provides unmarshaler without descriptors. It should help you if you have unknown/broken binary and hope to search correct descriptors from seeing actual values.

## How to use

- Just call `Unmarshal()`

```go
	msg := protodump.Message{}
	err = protodump.Unmarshal(bin, msg)
	if err != nil {
		log.Fatal(err)
	}
```

## CLI

- Extract a binary encoded in protobuf without descriotors as JSON.

```sh
$ echo -n "\x12\x07\x74\x65\x73\x74\x69\x6e\x67" | protodump
{"2":{"__bytes":"dGVzdGluZw==","__packed":[116,101,115,116,105,110,103],"__string":"testing"}}
```

- So you can extract specific fields with `jq`

```sh
$ echo -n "\x08\x01\x12\x07\x74\x65\x73\x74\x69\x6e\x67\x19\x0b\x00\x00\x00\x00\x00\x00\x00\x25\x6f\x00\x00\x00\x28\x01\x32\x02\x08\x01\xaa\x06\x02\x02\x03\xb2\x06\x03\x61\x61\x61\xb2\x06\x03\x62\x62\x62\xca\x06\x02\x00\x01\xd2\x06\x02\x08\x02\xd2\x06\x02\x08\x03" | protodump
{"1":1,"101":{"__bytes":"AgM=","__packed":[2,3],"__string":"\u0002\u0003"},"102":[{"__bytes":"YWFh","__packed":[97,97,97],"__string":"aaa"},{"__bytes":"YmJi","__packed":[98,98,98],"__string":"bbb"}],"105":{"__bytes":"AAE=","__packed":[0,1],"__string":"\u0000\u0001"},"106":[{"__bytes":"CAI=","__message":{"1":2},"__packed":[8,2],"__string":"\u0008\u0002"},{"__bytes":"CAM=","__message":{"1":3},"__packed":[8,3],"__string":"\u0008\u0003"}],"2":{"__bytes":"dGVzdGluZw==","__packed":[116,101,115,116,105,110,103],"__string":"testing"},"3":11,"4":111,"5":1,"6":{"__bytes":"CAE=","__message":{"1":1},"__packed":[8,1],"__string":"\u0008\u0001"}}
$ echo -n "\x08\x01\x12\x07\x74\x65\x73\x74\x69\x6e\x67\x19\x0b\x00\x00\x00\x00\x00\x00\x00\x25\x6f\x00\x00\x00\x28\x01\x32\x02\x08\x01\xaa\x06\x02\x02\x03\xb2\x06\x03\x61\x61\x61\xb2\x06\x03\x62\x62\x62\xca\x06\x02\x00\x01\xd2\x06\x02\x08\x02\xd2\x06\x02\x08\x03" | protodump | jq '."102"[].__string'
"aaa"
"bbb"
```

## TODO

- Support map type

## Limitation

- `repeated fixed??` types will not be recognized correctly
- Don't check each types in repeated fields
