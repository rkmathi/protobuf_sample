// $ go run go/read_sample.go

package main

import (
	"log"
	"io/ioutil"
	"github.com/golang/protobuf/proto"
	"github.com/rkmathi/samples"
	"os"
)

func main() {
	file, _ := os.Open("sample_output_data.pb")
	bytes, _ := ioutil.ReadAll(file)

	person := &person.Person {}
	err3 := proto.Unmarshal(bytes, person)
	if err3 != nil {
		log.Fatal("unmarshaling error: ", err3)
	}

	log.Printf("Unmarshalled to: %+v", person)
}
