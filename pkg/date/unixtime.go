package date

import (
	"strconv"
	"time"
)

func (formatter unixTimeInputFormatter) ToTime() (time.Time, error) {
	unixtime, _ := strconv.ParseInt(formatter.input, 10, 64)
	now := time.Now().Unix()
	if unixtime <= now {
		return time.Unix(unixtime, 0), nil
	}

	return time.Unix(0, unixtime), nil
}

func (formatter unixTimeInputFormatter) ToTimeStruct(t time.Time) TimeStruct {
	return newTimeStruct(formatter.input, t)
}
