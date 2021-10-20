package data

import (
	"encoding/json"
	"time"
)

type JsonTime struct {
	time.Time
}

func (jt *JsonTime) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	var err error
	jt.Time, err = time.Parse("2006-01-02T15:04", s)
	return err
}
