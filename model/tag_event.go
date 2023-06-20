package model

import (
	"cli/epc"
	"time"
)

type TagEvent struct {
	Timestamp         time.Time         `json:"timestamp"`
	Hostname          string            `json:"hostname"`
	EventType         string            `json:"eventType"`
	TagInventoryEvent TagInventoryEvent `json:"tagInventoryEvent"`
}

type TagInventoryEvent struct {
	Epc         string `json:"epc"`
	EpcHex      string `json:"epcHex"`
	AntennaPort int    `json:"antennaPort"`
	AntennaName string `json:"antennaName"`
}

func NewTagEvent(epc epc.EPC) TagEvent {
	b64, err := epc.B64()
	if err != nil {
		b64 = ""
	}

	return TagEvent{
		Timestamp: time.Now().In(time.UTC),
		Hostname:  "r700-emulator",
		EventType: "tagInventory",
		TagInventoryEvent: TagInventoryEvent{
			Epc:         b64,
			EpcHex:      epc.Hex(),
			AntennaPort: 1,
			AntennaName: "Antenna 1",
		},
	}
}
