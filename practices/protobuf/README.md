# 学习proto_buf

## simple sample

As you can see, each field in the message definition has a unique number. These numbers are used to identify your fields
in the message binary format, and should not be changed once your message type is in use.

```protobuf
message SearchRequest {
  required string query = 1;
  optional int32 page_number = 2;
  optional int32 result_per_page = 3;
}
```
