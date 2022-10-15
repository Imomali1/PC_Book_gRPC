package serializer

import (
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

// ProtoBufToJSON converts protocol buffer message to JSON string
func ProtoBufToJSON(message proto.Message) (string, error) {
	marshaller := jsonpb.Marshaler{
		OrigName:     true,
		EnumsAsInts:  false,
		EmitDefaults: true,
		Indent:       " ",
	}
	return marshaller.MarshalToString(message)
}
