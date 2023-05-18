package entity

import (
	"fmt"
	"strings"
	"time"
)

type Event struct {
	UserID int        `json:"user_id"`
	Name   string     `json:"name"`
	Date   CustomDate `json:"date"`
}

type CustomDate struct {
	time.Time
}

const layout = "2006-01-02 15:04:05"

func (c *CustomDate) UnmarshalJSON(b []byte) (err error) {
	s := strings.Trim(string(b), `"`)
	if s == "null" {
		return
	}
	c.Time, err = time.Parse(layout, s)
	return
}

func (c *CustomDate) MarshalJSON() ([]byte, error) {
	if c.Time.IsZero() {
		return nil, nil
	}
	return []byte(fmt.Sprintf(`"%s"`, c.Time.Format(layout))), nil
}
