package report

import (
	"errors"
	"strings"
	"time"
)

func FromCSVLineToTornadoMsg(line []byte) (TornadoMsg, error) {
	words := strings.Split(string(line), ",")
	if len(words) < 8 {
		return TornadoMsg{}, errors.New("line did not contain at least 8 columns")
	}
	distance, direction, location := GetDistanceFromLocation(words[2])

	return TornadoMsg{
		Time:      StringToUnixTime(time.Now().UTC().Format(time.DateOnly), words[0]),
		F_Scale:   StringToInt32orZero(words[1]),
		Distance:  distance,
		Direction: direction,
		Location:  location,
		County:    words[3],
		State:     words[4],
		Lat:       StringToInt32orZero(words[5]),
		Lon:       StringToInt32orZero(words[6]),
		Remarks:   words[7],
	}, nil
}
