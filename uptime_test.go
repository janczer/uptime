package main

import (
	"testing"
)

func TestSecondsToTime(t *testing.T) {
	test := map[int]map[string]int{
		15718:  {"hour": 4, "minut": 21, "sec": 58},
		216000: {"hour": 60, "minut": 0, "sec": 0},
		5552:   {"hour": 1, "minut": 32, "sec": 32},
	}
	for i, v := range test {
		h, m, s := secondsToTime(i)
		if h != v["hour"] {
			t.Errorf("Test faild. Hours: %d != %d", h, v["hour"])
		}
		if m != v["minut"] {
			t.Errorf("Test faild. Minutes: %d != %d", m, v["minut"])
		}
		if s != v["sec"] {
			t.Errorf("Test faild. Seconds: %d != %d", s, v["sec"])
		}
	}
}
