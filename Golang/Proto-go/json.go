package main

import (
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func toJSON(pb proto.Message) string {
	option := protojson.MarshalOptions{
		Multiline: true,
	}

	out, err := option.Marshal(pb)
	if err != nil {
		log.Fatalln("Failed to marshal:", err)
	}
	return string(out)
}

func fromJSON(json string, pb proto.Message) {
	option := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}

	err := option.Unmarshal([]byte(json), pb)
	if err != nil {
		log.Fatalln("Failed to unmarshal:", err)
	}
}
