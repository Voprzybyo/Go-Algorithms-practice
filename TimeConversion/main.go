package main

import (
	"fmt"
	"strconv"
	"strings"
)

func formatTime(s string) string {

	s = strings.Replace(s, "PM", "", 1)
	ss := strings.Split(s, ":")

	hour, _ := strconv.ParseInt(ss[0], 10, 0)
	hour = (hour + 12) % 24
	ss[0] = strconv.FormatInt(hour, 10)

	if ss[0] == "0" {
		ss[0] = "00"
	}

	ret := strings.Join(ss, ":")

	return ret
}

func convertTime(s string) string {

	var convertedTime string

	ss := strings.Split(s, ":")
	conversionNotNeeded := ((strings.Contains(s, "AM") && ss[0] != "12") ||
		(strings.Contains(s, "PM") && ss[0] == "12"))

	s = strings.Replace(s, "AM", "", 1)
	s = strings.Replace(s, "PM", "", 1)

	if conversionNotNeeded {
		return s
	} else {
		convertedTime = formatTime(s)
	}

	return convertedTime
}

func main() {

	s := "12:01:00PM"
	convertedTime := convertTime(s)
	fmt.Println(convertedTime)

}
