# protosl

A protobuf schema-less module.


## CLI

- Extract a binary encoded in protobuf without descriotors as JSON.

```sh
$ echo -n "\x12\x07\x74\x65\x73\x74\x69\x6e\x67" | base64 | protosl
{"2":{"__bytes":"dGVzdGluZw==","__packed":[116,101,115,116,105,110,103],"__string":"testing"}}
```

## TODO

- Support map type

## Limitation

- `repeated fixed??` types will not be recognized correctly
