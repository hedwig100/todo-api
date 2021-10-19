package data

import (
	"time"
)

type JsonTime struct {
	time.Time
}

func (jt *JsonTime) UnmarshalJSON(data []byte) (err error) {
	sdata := string(data)
	if sdata == "null" {
		return nil
	}

	jt.Time, err = time.Parse("2006-01-02T15:04", sdata)
	return
}
