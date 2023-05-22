package tools

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

func NewEPC(header uint8, manager uint32, class uint32, serial uint64) (EPC, error) {
	if header <= 0 {
		header = DefaultHeader
	}
	if manager <= 0 {
		manager = DefaultManager
	}
	if class <= 0 {
		class = uint32(rand.Int31n(99999999))
	}
	if serial <= 0 {
		serial = uint64(rand.Int63n(99999999999))
	}
	return EPC{
		Header:  0x35,
		Manager: manager,
		Class:   class,
		Serial:  serial,
	}, nil
}
