package models

import "time"

type Endpoint struct {
	URL       string `json:"url" bson:"url"`
	Timestamp int64  `json:"timestamp" bson:"timestamp"`
}

func (e *Endpoint) SetTimestamp() int64 {
	e.Timestamp = time.Now().Unix()
	return e.Timestamp
}
