package epc

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"math/rand"
	"strconv"
)

const (
	DefaultHeader  = 0x35
	DefaultManager = 759936
	MaxHeader      = 0xFF
	MaxManager     = 0xFFFFFFF
	MaxClass       = 0xFFFFFF
	MaxSerial      = 0xFFFFFFFFF
)

type EPC struct {
	Header  uint8
	Manager uint32
	Class   uint32
	Serial  uint64
}

func (epc EPC) PureIdentityURI() string {
	return fmt.Sprintf("urn:epc:id:gid:%d.%d.%d", epc.Manager, epc.Class, epc.Serial)
}

func (epc EPC) TagURI() string {
	return fmt.Sprintf("urn:epc:tag:gid-96:%d.%d.%d", epc.Manager, epc.Class, epc.Serial)
}

func (epc EPC) Hex() string {
	return fmt.Sprintf("%02X%07X%06X%09X", epc.Header, epc.Manager, epc.Class, epc.Serial)
}

func (epc EPC) B64() (string, error) {
	hex := epc.Hex()
	buf := make([]byte, 12)
	for i := 0; i < 12; i++ {
		v, err := strconv.ParseInt(hex[i*2:(i*2)+2], 16, 0)
		if err != nil {
			return "", err
		}
		buf[i] = byte(v)
	}
	b64 := bytes.Buffer{}
	_, err := base64.NewEncoder(base64.StdEncoding, &b64).Write(buf)
	if err != nil {
		return "", err
	}

	return b64.String(), nil
}

func NewEPC(header uint8, manager uint32, class uint32, serial uint64) (*EPC, error) {
	if header <= 0 {
		header = DefaultHeader
	}
	if header > MaxHeader {
		return nil, fmt.Errorf("invalid header value")
	}

	if manager <= 0 {
		manager = DefaultManager
	}
	if manager > MaxManager {
		return nil, fmt.Errorf("invalid manager value")
	}

	if class <= 0 {
		class = uint32(rand.Int31n(MaxClass))
	}
	if class > MaxClass {
		return nil, fmt.Errorf("invalid class value")
	}

	if serial <= 0 {
		serial = uint64(rand.Int63n(MaxSerial))
	}
	if serial > MaxSerial {
		return nil, fmt.Errorf("invalid serial value")
	}

	return &EPC{
		Header:  0x35,
		Manager: manager,
		Class:   class,
		Serial:  serial,
	}, nil
}

func DecodeHex(hex string) (*EPC, error) {
	if len(hex) != 24 {
		return nil, fmt.Errorf("invalid EPC hex length")
	}

	header, err := strconv.ParseUint(hex[0:2], 16, 8)
	if err != nil {
		return nil, err
	}

	manager, err := strconv.ParseUint(hex[2:9], 16, 32)
	if err != nil {
		return nil, err
	}

	class, err := strconv.ParseUint(hex[9:15], 16, 32)
	if err != nil {
		return nil, err
	}

	serial, err := strconv.ParseUint(hex[15:], 16, 64)
	if err != nil {
		return nil, err
	}

	return &EPC{
		Header:  uint8(header),
		Manager: uint32(manager),
		Class:   uint32(class),
		Serial:  serial,
	}, nil
}

func DecodeB64(b64 string) (*EPC, error) {
	if len(b64) != 16 {
		return nil, fmt.Errorf("invalid EPC base64 length")
	}
	buf := make([]byte, 12)
	_, err := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(b64)).Read(buf)
	if err != nil {
		return nil, err
	}

	var hex string
	for _, b := range buf {
		hex += fmt.Sprintf("%02X", b)
	}

	return DecodeHex(hex)
}
