package report

import (
	"fmt"
)

//go:generate stringer -type=ReportType
type ReportType int

const (
	Hail ReportType = iota
	Wind
	Tornado
)

func FromString(s string) (ReportType, error) {
	switch s {
	case "Hail":
		return Hail, nil
	case "Wind":
		return Wind, nil
	case "Tornado":
		return Tornado, nil
	}
	var r ReportType
	return r, fmt.Errorf("%q does not map to a known ReportType", s)
}
