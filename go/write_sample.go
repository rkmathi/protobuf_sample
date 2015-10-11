// $ go run go/write_sample.go

package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"github.com/golang/protobuf/proto"
	"github.com/rkmathi/samples"
	"os"
)

func main() {
	person := &person.Person {
		Id: proto.Int32(789),
		Name: proto.String("NAME from Go"),
		Email: proto.String("go@example.com"),
	}

	data, err1 := proto.Marshal(person)
	if err1 != nil {
		log.Fatal("marshaling error: ", err1)
	}
	//log.Printf("marshalled to: %+v", data)

	buf := new(bytes.Buffer)
	err2 := binary.Write(buf, binary.LittleEndian, data)
	if err2 != nil {
		log.Fatal("binary.Write error: ", err2)
	}

	file, _ := os.Create("sample_output_data.pb")
	defer file.Close()
	file.Write(buf.Bytes())
}
