package dto

import (
	"encoding/json"
	"time"
)

type TimeDTO struct {
	Time time.Time
}

func (t TimeDTO) MarshalJSON() ([]byte, error) {
	formattedTime := t.Time.Format("2006-01-02 15:04:05")
	return json.Marshal(formattedTime)
}
