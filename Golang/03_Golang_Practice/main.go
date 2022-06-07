package main

import (
	"ProtoBufTest/src/complexpb"
	"ProtoBufTest/src/enumpb"
	"ProtoBufTest/src/simplepb"
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func main() {

	sm := doSimple()
	readAndWriteDemo(sm)
	jsonDemo(sm)
	doEnum()
	doComplex()

}

func doComplex() {

	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First Message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   1,
				Name: "Second Message",
			},
			&complexpb.DummyMessage{
				Id:   1,
				Name: "Third Message",
			},
		},
	}

	fmt.Println(cm)
}

func doEnum() {

	em := enumpb.EnumMessage{
		Id:        42,
		DayOfWeek: enumpb.DayOfWeek_THURSDAY,
	}

	em.DayOfWeek = enumpb.DayOfWeek_MONDAY
	fmt.Println(em)
}

func jsonDemo(sm proto.Message) {
	smAsString := toJSON(sm)
	sm2 := &simplepb.SimplePerson{}
	fromJSON(smAsString, sm2)
	fmt.Println("From JSON :", sm2)
}

func fromJSON(in string, pb proto.Message) {

	err := protojson.Unmarshal([]byte(in), pb)
	if err != nil {
		log.Fatalln("Error unmarshaling string to PB Struct")
	}
}

func toJSON(pb proto.Message) string {

	marshaler := protojson.MarshalOptions{}
	out, err := marshaler.Marshal(pb)
	if err != nil {
		log.Fatalln("Error marshaling message to json.")
		return ""
	}
	return string(out)
}

func readAndWriteDemo(sm proto.Message) {

	writeToFile("SimpleTest.bin", sm)
	sm2 := &simplepb.SimplePerson{}
	readFromFile("SimpleTest.bin", sm2)
	fmt.Println("From File :", sm2)
}

func writeToFile(fname string, pb proto.Message) error {

	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialise in bytes", err)
		return err
	}
	if ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}
	return nil
}

func readFromFile(fname string, pb proto.Message) error {

	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Error reading binary file")
		return err
	}

	err = proto.Unmarshal(in, pb)
	if err != nil {
		log.Fatalln("Error during unmarshal")
	}
	return nil
}

func doSimple() *simplepb.SimplePerson {
	sm := simplepb.SimplePerson{
		Age:               29,
		FirstName:         "Sheryar",
		LastName:          "Butt",
		IsProfileVerified: true,
		Height:            172,
	}

	return &sm
}
