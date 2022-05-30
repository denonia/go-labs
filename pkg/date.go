package pkg

import (
	"fmt"
	"strings"
	"time"
)

const dateFormat = "2006-01-02"

type JSONDate struct {
	time.Time
}

// interface Marshaler
func (d *JSONDate) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(d.Time).Format(dateFormat))
	return []byte(stamp), nil
}

// interface Unmarshaler
func (d *JSONDate) UnmarshalJSON(data []byte) (err error) {
	s := strings.Trim(string(data), "\"")
	if s == "null" {
		d.Time = time.Time{}
		return
	}
	d.Time, err = time.Parse(dateFormat, s)
	return
}
