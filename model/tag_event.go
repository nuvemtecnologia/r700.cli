package model

import "time"

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

func NewTagEvent(epc string) TagEvent {
	return TagEvent{
		Timestamp: time.Now(),
		EventType: "TagInventoryEvent",
		TagInventoryEvent: TagInventoryEvent{
			Epc:         epc,
			EpcHex:      epc,
			AntennaPort: 1,
			AntennaName: "Antenna 1",
		},
	}
}
