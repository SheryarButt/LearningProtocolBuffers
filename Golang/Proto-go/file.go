package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/proto"
)

func writeToFile(fileName string, pb proto.Message) {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Failed to encode proto:", err)
		return
	}

	if err := ioutil.WriteFile(fileName, out, 0644); err != nil {
		log.Fatalf("can't write file: %v", err)
		return
	}

	fmt.Println("Data written to", fileName)
}

func readFromFile(fileName string, pb proto.Message) {
	in, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Failed to read file:", err)
		return
	}

	if err := proto.Unmarshal(in, pb); err != nil {
		log.Fatalln("Failed to parse file:", err)
		return
	}

	fmt.Println("Data read from", fileName)
}
