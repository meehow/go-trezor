package trezor

import (
	"context"
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/meehow/go-trezor/pb"
	"github.com/trezor/trezord-go/core"
	"github.com/trezor/trezord-go/memorywriter"
	"github.com/trezor/trezord-go/usb"
	"github.com/trezor/trezord-go/wire"
	"google.golang.org/protobuf/proto"
)

const attempts = 20

type Trezor struct {
	core       *core.Core
	path       string
	ssid       string
	passphrase string
	logger     *log.Logger
	debug      bool
}

type Config struct {
	DeriveCardano bool
	Logger        *log.Logger
	Passphrase    string
	EmulatorPort  int
	Debug         bool
}

type Message struct {
	Kind pb.MessageType
	Data []byte
}

func New(cfg Config) (*Trezor, error) {
	longMemoryWriter := memorywriter.New(90000, 200, true, nil)
	libusb, err := usb.InitLibUSB(longMemoryWriter, true, true, true)
	if err != nil {
		return nil, fmt.Errorf("libusb: %w", err)
	}
	bus := []core.USBBus{libusb}
	if cfg.EmulatorPort != 0 {
		udp, err := usb.InitUDP([]usb.PortTouple{{
			Normal: cfg.EmulatorPort,
			Debug:  cfg.EmulatorPort + 1,
		}}, longMemoryWriter)
		if err != nil {
			return nil, err
		}
		bus = append(bus, udp)
	}
	b := usb.Init(bus...)
	c := core.New(b, longMemoryWriter, true, true)
	var enums []core.EnumerateEntry
	for i := 1; i <= attempts; i++ {
		enums, err = c.Enumerate()
		if err != nil {
			return nil, err
		}
		if len(enums) > 0 {
			err = nil
			break
		}
		if i == attempts {
			return nil, errors.New("trezor not found")
		}
		fmt.Println("connect trezor")
		time.Sleep(time.Second * 3)
	}
	fmt.Printf("%+v\n", enums[0])
	tr := &Trezor{
		core:       c,
		path:       enums[0].Path,
		passphrase: cfg.Passphrase,
		logger:     cfg.Logger,
		debug:      cfg.Debug,
	}
	if tr.logger == nil {
		tr.logger = log.New(io.Discard, "", 0)
	}
	if enums[0].Session == nil {
		tr.ssid, err = c.Acquire(tr.path, "", cfg.Debug)
		if err != nil {
			return nil, err
		}
	} else {
		tr.ssid = *enums[0].Session
	}
	fmt.Println("ssid", tr.ssid)
	if !cfg.Debug {
		sessionID := make([]byte, 32)
		rand.Read(sessionID)
		_, err = tr.Call(&pb.Initialize{
			DeriveCardano: &cfg.DeriveCardano,
			SessionId:     sessionID,
		})
	}
	return tr, err
}

func (tr *Trezor) Call(msg proto.Message) (*Message, error) {
	binbody, err := Encode(msg)
	if err != nil {
		return nil, err
	}
	switch msg.(type) {
	case *pb.PassphraseAck, *pb.LoadDevice, *pb.FirmwareUpload:
		tr.logger.Printf("---> MessageType_%v", msg.ProtoReflect().Type().Descriptor().Name())
	default:
		tr.logger.Printf("---> MessageType_%v %x", msg.ProtoReflect().Type().Descriptor().Name(), binbody)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()
	binres, err := tr.core.Call(binbody, tr.ssid, core.CallModeReadWrite, tr.debug, ctx)
	if err != nil {
		return nil, err
	}
	resp, err := Decode(binres)
	if err != nil {
		return nil, err
	}

	tr.logger.Printf("<--- %v %x", resp.Kind.String(), binres)

	switch resp.Kind {
	case pb.MessageType_MessageType_Failure:
		msg := new(pb.Failure)
		err = proto.Unmarshal(resp.Data, msg)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("trezor: %s", msg.GetMessage())
	case pb.MessageType_MessageType_ButtonRequest:
		msg := new(pb.ButtonRequest)
		err = proto.Unmarshal(resp.Data, msg)
		if err != nil {
			return nil, err
		}
		tr.logger.Println("ButtonRequest", msg.String())
		return tr.Call(&pb.ButtonAck{})
	case pb.MessageType_MessageType_PassphraseRequest:
		return tr.Call(&pb.PassphraseAck{
			Passphrase: &tr.passphrase,
		})
	}
	return Decode(binres)
}

func Encode(msg proto.Message) ([]byte, error) {
	name := "MessageType_" + string(msg.ProtoReflect().Type().Descriptor().Name())
	kind := uint16(pb.MessageType_value[name])
	data, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}
	var header [6]byte
	binary.BigEndian.PutUint16(header[0:2], kind)
	binary.BigEndian.PutUint32(header[2:6], uint32(len(data)))
	return append(header[:], data...), nil
}

func Decode(binbody []byte) (*Message, error) {
	if len(binbody) < 6 {
		return nil, core.ErrMalformedData
	}
	kind := binary.BigEndian.Uint16(binbody[0:2])
	size := binary.BigEndian.Uint32(binbody[2:6])
	data := binbody[6:]
	if uint32(len(data)) != size {
		return nil, core.ErrMalformedData
	}
	if wire.Validate(data) != nil {
		return nil, core.ErrMalformedData
	}
	return &Message{
		Kind: pb.MessageType(kind),
		Data: data,
	}, nil
}
