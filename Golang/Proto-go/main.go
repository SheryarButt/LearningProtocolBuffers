package main

import (
	pb "Proto-go/proto"
	"fmt"
	"reflect"

	"google.golang.org/protobuf/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          1,
		IsSimple:    true,
		Name:        "Simple",
		SampleLists: []int32{1, 2, 3},
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy{
			Id:   1,
			Name: "Dummy",
		},
		MultipleDummies: []*pb.Dummy{
			{Id: 1, Name: "Dummy1"},
			{Id: 2, Name: "Dummy2"},
			{Id: 3, Name: "Dummy3"},
		},
	}
}

func doEnum() *pb.Enumeration {
	return &pb.Enumeration{
		EyeColor: pb.EyeColor_EYE_COLOR_BLUE,
	}
}

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IdWrapper{
			"myid":  {Id: 42},
			"myid2": {Id: 43},
			"myid3": {Id: 44},
		},
	}
}

func doOneOf(message interface{}) interface{} {
	switch X := message.(type) {
	case *pb.Result_Id:
		return (message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		return (message.(*pb.Result_Message).Message)
	default:
		return fmt.Sprintf("Unknown message type: %T", X)
	}
}

func doFile(p proto.Message) interface{} {
	path := "simple.bin"
	writeToFile(path, p)

	message := &pb.Simple{}
	readFromFile(path, message)
	return message
}

func doToJSON(p proto.Message) string {
	return toJSON(p)
}

func doFromJSON(jsonString string, t reflect.Type) proto.Message {

	message := reflect.New(t.Elem()).Interface().(proto.Message)
	fromJSON(jsonString, message)
	return message
}

func main() {
	fmt.Println("Simple Message : ", doSimple())
	fmt.Println("Complex Message : ", doComplex())
	fmt.Println("Enumeration : ", doEnum())
	fmt.Println("Map Values : ", doMap())
	fmt.Println("ID set to : ", doOneOf(&pb.Result_Id{Id: 1}))
	fmt.Println("Message set to :", doOneOf(&pb.Result_Message{Message: "Hello"}))
	fmt.Println("File : ", doFile(doSimple()))

	simpleJSON := doToJSON(doSimple())
	fmt.Println("simpleJSON : ", simpleJSON)
	fmt.Println("simpleJSON : ", doFromJSON(simpleJSON, reflect.TypeOf(&pb.Simple{})))

	complexJSON := doToJSON(doComplex())
	fmt.Println("complexJSON : ", complexJSON)
	fmt.Println("complexJSON : ", doFromJSON(complexJSON, reflect.TypeOf(&pb.Complex{})))

	fmt.Println("Unmarshal Option : ", doFromJSON(`{"id":1,"unknown":"Dummy"}`, reflect.TypeOf(&pb.Simple{})))
}
