package mbc

import (
	"encoding/json"
	"fmt"
	"time"
)

// UnixTime represents an UnixTime
type UnixTime struct {
	time.Time
}

func (u UnixTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", u.Unix())), nil
}

func (u *UnixTime) UnmarshalJSON(bytes []byte) error {
	var unix int64
	if err := json.Unmarshal(bytes, &unix); err != nil {
		return err
	}
	u.Time = time.Unix(unix, 0)
	return nil
}

type Log interface {
	Infof(fmt string, params ...interface{})
	Debugf(fmt string, params ...interface{})
	Warningf(fmt string, params ...interface{})
	Errorf(fmt string, params ...interface{})
}

type logMock struct{}

func (l logMock) Infof(fmt string, params ...interface{}) {}

func (l logMock) Debugf(fmt string, params ...interface{}) {}

func (l logMock) Warningf(fmt string, params ...interface{}) {}

func (l logMock) Errorf(fmt string, params ...interface{}) {}
