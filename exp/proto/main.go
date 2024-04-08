package main

import (
	"fmt"
	"log"

	"github.com/moroz/saboru-server/exp/proto/pb"
	"google.golang.org/protobuf/proto"
)

func main() {
	channel := pb.Channel{
		Id:    42,
		Label: "general",
	}

	out, err := proto.Marshal(&channel)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", out)

	decoded := &pb.Channel{}
	err = proto.Unmarshal(out, decoded)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", decoded)
}
