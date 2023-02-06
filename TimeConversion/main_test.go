package main

import "testing"

func TestConvertTime(t *testing.T) {

	type test struct {
		inputTime  string
		outputTime string
	}

	tests := []test{
		{"07:05:45PM", "19:05:45"},
		{"12:01:00PM", "12:01:00"},
		{"12:01:00AM", "00:01:00"},
		{"01:15:59PM", "13:15:59"},
		{"02:12:35AM", "02:12:35"},
		{"08:13:13PM", "20:13:13"},
		{"09:37:00PM", "21:37:00"},
	}

	for _, v := range tests {
		s := convertTime(v.inputTime)

		if s != v.outputTime {
			t.Error("Expected:", v.outputTime, "Got:", s)
		}
	}
}
