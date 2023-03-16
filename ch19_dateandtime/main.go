package main

import (
	"fmt"
	"time"
)

// PrintTime print time )
func PrintTime(label string, t *time.Time) {
	/*Printfln(
		"%s: Day: %v: Month: %v Year: %v",
		label, t.Day(), t.Month(), t.Year(),
	)
	//Format
	layout := "Day: 02 Month: Jan Year: 2006" //reference time 2 January 2006 03:04:05 PM in the time zone UTC -7
	fmt.Println("layout", label, t.Format(layout))*/
	fmt.Println("layout RFC822Z", label, t.Format(time.RFC822Z))
}

func main() {
	Printfln("Hello, Dates and Times")

	current := time.Now()
	specific := time.Date(1995, time.June, 9, 0, 0, 0, 0, time.Local)
	unix := time.Unix(1433228090, 0)
	PrintTime("Current", &current)
	PrintTime("Specific", &specific)
	PrintTime("UNIX", &unix)

	//Parse
	layout := "2006-Jan-02 15:04:05" //reference time 2 January 2006 03:04:05 PM in the time zone UTC -7
	dates := []string{
		"1995-Jun-09 04:00:33",
		"2015-Jun-02 14:03:44",
	}
	for _, d := range dates {
		time, err := time.Parse(layout, d)
		if err == nil {
			PrintTime("Parsed", &time)
		} else {
			Printfln("Error: %s", err.Error())
		}
	}
	//rfc
	dates2 := []string{
		"09 Jun 95 00:00 GMT",
		"02 Jun 15 00:00 GMT",
	}
	for _, d := range dates2 {
		time, err := time.Parse(time.RFC822, d)
		if err == nil {
			PrintTime("Parsed RFC822", &time)
		} else {
			Printfln("Error RFC822: %s", err.Error())
		}
	}

	//Location
	layoutLocation := "02 Jan 06 15:04"
	date := "09 Jun 95 19:30"
	london, lonerr := time.LoadLocation("Europe/London")
	newyork, nycerr := time.LoadLocation("America/New_York")
	local, _ := time.LoadLocation("Local")
	local2 := time.FixedZone("EDT", -4*60*60)
	if lonerr == nil && nycerr == nil {
		nolocation, _ := time.Parse(layoutLocation, date)
		londonTime, _ := time.ParseInLocation(layoutLocation, date, london)
		newyorkTime, _ := time.ParseInLocation(layoutLocation, date, newyork)
		localTime, _ := time.ParseInLocation(layoutLocation, date, local)
		localTime2, _ := time.ParseInLocation(layoutLocation, date, local2)
		PrintTime("No location:", &nolocation)
		PrintTime("London:", &londonTime)
		PrintTime("New York:", &newyorkTime)
		PrintTime("Local:", &localTime)
		PrintTime("Local2:", &localTime2)
	} else {
		fmt.Println(lonerr.Error(), nycerr.Error())
	}
	//control
	t, err := time.Parse(time.RFC822, "09 Jun 95 04:59 BST")
	if err == nil {
		Printfln("After: %v", t.After(time.Now()))
		Printfln("Round: %v", t.Round(time.Hour))
		Printfln("Truncate: %v", t.Truncate(time.Hour))
	} else {
		fmt.Println(err.Error())
	}
	//equal
	t1, _ := time.Parse(time.RFC822Z, "09 Jun 95 04:59 +0100")
	t2, _ := time.Parse(time.RFC822Z, "08 Jun 95 23:59 -0400")
	Printfln("Equal Method: %v", t1.Equal(t2))
	Printfln("Equality Operator: %v", t1 == t2)

	//Duration
	var d time.Duration = time.Hour + (30 * time.Minute)
	Printfln("Hours: %v", d.Hours())
	Printfln("Mins: %v", d.Minutes())
	Printfln("Seconds: %v", d.Seconds())
	Printfln("Millseconds: %v", d.Milliseconds())
	rounded := d.Round(time.Hour)
	Printfln("Rounded Hours: %v", rounded.Hours())
	Printfln("Rounded Mins: %v", rounded.Minutes())
	trunc := d.Truncate(time.Hour)
	Printfln("Truncated Hours: %v", trunc.Hours())
	Printfln("Rounded Mins: %v", trunc.Minutes())

	//between
	toYears := func(d time.Duration) int {
		return int(d.Hours() / (24 * 365))
	}
	future := time.Date(2051, 0, 0, 0, 0, 0, 0, time.Local)
	past := time.Date(1965, 0, 0, 0, 0, 0, 0, time.Local)
	Printfln("Future: %v", toYears(time.Until(future)))
	Printfln("Past: %v", toYears(time.Since(past)))

	//ParseDuration
	pD, err := time.ParseDuration("1h30m")
	if err == nil {
		Printfln("Hours: %v", pD.Hours())
		Printfln("Mins: %v", pD.Minutes())
		Printfln("Seconds: %v", pD.Seconds())
		Printfln("Millseconds: %v", pD.Milliseconds())
	} else {
		fmt.Println(err.Error())
	}

}
