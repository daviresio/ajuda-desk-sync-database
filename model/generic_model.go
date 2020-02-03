package model

import "encoding/json"

type GerenicData struct{
	Payload struct{
		Before *struct{} `json:"before"`
		After json.RawMessage `json:"after"`
		Source struct{
			Table string `json:"table"`
		}
	} `json:"payload"`

}

type IdData struct {
	Payload struct{
		After struct{
			Id int `json:"id"`
		} `json:"after"`
	} `json:"payload"`
}