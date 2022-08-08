package utils

import (
	"encoding/json"
	"fmt"
	"time"
)

// UnixTime represents an UnixTime
type UnixTime time.Time

func (u UnixTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", time.Time(u).Unix())), nil
}

func (u *UnixTime) UnmarshalJSON(bytes []byte) error {
	var unix int64
	if err := json.Unmarshal(bytes, &unix); err != nil {
		return err
	}
	*u = UnixTime(time.Unix(unix, 0))
	return nil
}

type Log interface {
	Infof(fmt string, params ...interface{})
	Debugf(fmt string, params ...interface{})
	Warningf(fmt string, params ...interface{})
	Errorf(fmt string, params ...interface{})
}

type DummyLog struct{}

func (l DummyLog) Infof(fmt string, params ...interface{}) {}

func (l DummyLog) Debugf(fmt string, params ...interface{}) {}

func (l DummyLog) Warningf(fmt string, params ...interface{}) {}

func (l DummyLog) Errorf(fmt string, params ...interface{}) {}
