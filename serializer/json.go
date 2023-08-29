package serializer

import (
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func ProtobufToJSON(message proto.Message) (string, error) {
	marshaler := protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: true,
		Indent:          "  ",
		UseProtoNames:   true,
	}

	data, err := marshaler.Marshal(message)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
