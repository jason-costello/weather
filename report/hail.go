package report

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	directions = map[string]struct{}{
		"N":   {},
		"NNE": {},
		"NE":  {},
		"ENE": {},
		"E":   {},
		"ESE": {},
		"SE":  {},
		"SSE": {},
		"S":   {},
		"SSW": {},
		"SW":  {},
		"WSW": {},
		"W":   {},
		"WNW": {},
		"NW":  {},
		"NNW": {},
	}
)

func FromCSVLineToHailMsg(line []byte) (HailMsg, error) {
	words := strings.Split(string(line), ",")
	if len(words) < 8 {
		return HailMsg{}, errors.New("line did not contain at least 8 columns")
	}
	distance, direction, location := GetDistanceFromLocation(words[2])
	return HailMsg{
		Time:      StringToUnixTime(time.Now().UTC().Format(time.DateOnly), words[0]),
		Size:      StringToInt32orZero(words[1]),
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

func GetDistanceFromLocation(loc string) (int32, string, string) {
	var distance int32
	var direction, location string
	words := strings.Split(loc, " ")
	for i, word := range words {
		if n, err := strconv.Atoi(word); err == nil {
			distance = int32(n)
			continue
		}

		if _, ok := directions[strings.ToUpper(word)]; ok {
			direction = strings.ToUpper(word)
			continue
		}
		location = strings.Join(words[i:], " ")
		break
	}
	return distance, direction, location
}
func StringToInt32orZero(size string) int32 {
	i, err := strconv.Atoi(size)
	if err != nil {
		i = 0
	}
	return int32(i)
}
func StringToUnixTime(dateOnly string, hhmm string) int64 {
	if len(hhmm) < 4 {
		return 0
	}

	if _, err := time.Parse(time.DateOnly, dateOnly); err != nil {
		return 0
	}

	t := fmt.Sprintf("%s %s:%s:00", dateOnly, hhmm[0:2], hhmm[2:4])
	newTime, err := time.Parse(time.DateTime, t)
	if err != nil {
		return 0
	}
	return newTime.UTC().Unix()

}
