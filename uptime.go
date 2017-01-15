package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println(getUptime(), getLoadAvg())
}

//Read file /proc/uptime
//and return string like " 15:50:10 up 0:27, power on 15:22:54,"
func getUptime() string {
	dat, err := ioutil.ReadFile("/proc/uptime")
	if err != nil {
		return fmt.Sprintf("Error with open file /proc/uptime. Error: \"%v\"\n", err)
	}
	times := strings.Fields(string(dat))
	t, _ := strconv.ParseFloat(times[0], 64)

	hours, minutes, seconds := secondsToTime(int(t))
	tim := time.Now()
	onHour := tim.Hour() - hours
	onMinute := tim.Minute() - minutes
	onSecond := tim.Second() - seconds
	if onMinute < 0 {
		onHour--
		onMinute += 60
	}
	if onSecond < 0 {
		onMinute--
		onSecond += 60
	}

	return fmt.Sprintf(" %d:%02d:%02d up %d:%02d, power on %d:%02d:%02d,", tim.Hour(), tim.Minute(), tim.Second(), hours, minutes, onHour, onMinute, onSecond)
}

func secondsToTime(seconds int) (int, int, int) {
	second := seconds % 60
	seconds = (seconds - second) / 60

	minutes := seconds % 60
	seconds = (seconds - minutes) / 60

	hours := seconds / 60

	return hours, minutes, second
}

//Read file: /proc/loadavg
//return string like "load average: 0.14, 0.34, 0.37"
func getLoadAvg() string {
	dat, err := ioutil.ReadFile("/proc/loadavg")
	if err != nil {
		return fmt.Sprintf("\nError with open file /proc/loadavg. Error: \"%v\"", err)
	}
	loadavg := strings.Fields(string(dat))

	return fmt.Sprintf("load average: %s, %s, %s", loadavg[0], loadavg[1], loadavg[2])
}
