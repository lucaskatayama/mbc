package mbc

import (
	"fmt"
	"strings"
)

// InstrumentSymbol represents an instrument symbol with format BASE-QUOTE
// e.g. BTC-BRL
type InstrumentSymbol string

const (
	BTCBRL InstrumentSymbol = "BTC-BRL"
	LTCBRL InstrumentSymbol = "LTC-BRL"
)

func (i InstrumentSymbol) String() string {
	return strings.ToUpper(string(i))
}

// Base returns the base symbol
func (i InstrumentSymbol) Base() string {
	parts := strings.Split(strings.ToUpper(string(i)), "-")
	return parts[0]
}

// Quote returns the quote symbol
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

// AssetSymbol represents an asset symbol
// e.g BTC
type AssetSymbol string

func (a AssetSymbol) String() string {
	return strings.ToUpper(string(a))
}

func (a AssetSymbol) normalize() AssetSymbol {
	return AssetSymbol(a.String())
}
