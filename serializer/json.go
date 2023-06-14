package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// ProtobufToJSON converts protocol buffer message to JSON string
func ProtobufToJSON(message proto.Message) (string, error) {
	marsahaller := protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		Indent:          "  ",
		UseProtoNames:   true,
	}
	data, err := marsahaller.Marshal(message)

	return string(data), err
}
