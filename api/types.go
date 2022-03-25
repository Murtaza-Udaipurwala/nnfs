package api

import "time"

type MetaData struct {
	Expiry    time.Time `json:"expiry"`
	IsOneTime bool      `json:"is_one_time"`
}
