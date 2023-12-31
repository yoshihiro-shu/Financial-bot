package entity

import "time"

type Stocks struct {
	ID     int32     `json:"id"`
	Symbol string    `json:"symbol"`
	Name   string    `json:"name"`
	Open   float32   `json:"open"`
	Close  float32   `json:"close"`
	High   float32   `json:"high"`
	Low    float32   `json:"low"`
	Volume int32     `json:"volume"`
	Date   time.Time `json:"date"`
}
