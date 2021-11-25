package mbc

import (
	"fmt"
	"strings"
)

// InstrumentSymbol represents an instrument symbol with format BASE-QUOTE
type InstrumentSymbol string

const (
	BTCBRL InstrumentSymbol = "BTC-BRL"
	LTCBRL InstrumentSymbol = "LTC-BRL"
)

func (i InstrumentSymbol) String() string {
	return strings.ToUpper(string(i))
}

func (i InstrumentSymbol) Base() string {
	parts := strings.Split(strings.ToUpper(string(i)), "-")
	return parts[0]
}

func (i InstrumentSymbol) Quote() string {
	parts := strings.Split(strings.ToUpper(string(i)), "-")
	return parts[1]
}

func (i InstrumentSymbol) normalize() InstrumentSymbol {
	s := strings.ToUpper(string(i))
	if strings.Contains(s, "-") {
		return i
	}

	return InstrumentSymbol(fmt.Sprintf("%s-%s", s[3:], s[:3]))
}

func (i InstrumentSymbol) toMB() InstrumentSymbol {
	return InstrumentSymbol(fmt.Sprintf("%s%s", i.Quote(), i.Base()))
}
