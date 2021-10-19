package data

import (
	"time"
)

type jsonTime struct {
	JsonTime time.Time
}

func (jt *jsonTime) UnmarshalJSON(data []byte) (err error) {
	sdata := string(data)
	if sdata == "null" {
		return nil
	}

	jt.JsonTime, err = time.Parse("2006-01-02T15:04", sdata)
	return
}
