# protosl

A protobuf schema-less module.


## CLI

- Extract a binary encoded in protobuf without descriotors as JSON.

```sh
$ echo -n "\x12\x07\x74\x65\x73\x74\x69\x6e\x67" | base64 | protosl
{"__2":"testing"}
```

## TODO

- Support repeated reputation type
- Support map type
