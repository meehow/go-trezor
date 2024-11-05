package trezor

import (
	"log"
	"testing"

	"github.com/meehow/go-trezor/pb"
	"google.golang.org/protobuf/proto"
)

func TestEmulator(t *testing.T) {
	tr, err := New(Config{
		Logger:       log.Default(),
		EmulatorPort: 21324,
	})
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

func TestDebug(t *testing.T) {
	tr, err := New(Config{
		Logger:       log.Default(),
		EmulatorPort: 21324,
		Debug:        true,
	})
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
	_, err = tr.Call(&pb.LoadDevice{
		Mnemonics: []string{"radar music civil monster device rare click radio basic turkey soft radar"},
	})
	if err != nil {
		t.Fatal(err)
	}
	// msg, err := tr.Call(&pb.GetFeatures{})
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// if msg.Kind != pb.MessageType_MessageType_Features {
	// 	t.Fatalf("unexpected message %v", msg.Kind)
	// }
	// features := new(pb.Features)
	// err = proto.Unmarshal(msg.Data, features)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// log.Printf("%+v", features)
}
