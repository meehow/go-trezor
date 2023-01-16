package trezor

import (
	"log"
	"testing"

	"github.com/meehow/go-trezor/pb"
	"google.golang.org/protobuf/proto"
)

func TestSimulator(t *testing.T) {
	tr, err := New(&Config{SimulatorPort: 21324})
	if err != nil {
		t.Fatal(err)
	}
	msg, err := tr.Call(&pb.GetFeatures{})
	if err != nil {
		t.Fatal(err)
	}
	if msg.Kind != pb.MessageType_MessageType_Features {
		t.Fatalf("unexpected message %v", msg.Kind)
	}
	features := new(pb.Features)
	err = proto.Unmarshal(msg.Data, features)
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("%+v", features)
}
