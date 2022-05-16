package main

import (
	"fmt"
	"time"
)

const (
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
)

func main() {
	now := time.Now()
	unix := time.Unix(100, 100)
	fmt.Println(now)
	fmt.Println(unix)
	fmt.Println(now.Format("Mon, Jan 2, 2006 at 3:04pm"))
	fmt.Println("Year: ", now.Year())
	fmt.Println("Month: ", now.Month())

	fmt.Println(now.Format(time.RFC850))
	fmt.Println(now.Format(time.RFC1123))

	est, _ := time.LoadLocation("EST")
	july4 := time.Date(1776, 7, 4, 12, 15, 0, 0, est)
	fmt.Println(july4)

	if july4.Before(now) {
		fmt.Println("July 4 is before Now ")
	}
	diff := now.Sub(july4)
	fmt.Printf("July 4 was about %f hours ago \n", diff.Hours())

	input_format := "1/2/2006 3:04pm"
	user_str := "4/16/2014 11:38am"
	user_date, err := time.Parse(input_format, user_str)
	if err != nil {
		fmt.Println(">>> Error parsing date string")
	}
	fmt.Println("User date = ", user_date.Format("Jan 2, 2006 @ 3:04pm"))

	twodaysDiff := time.Hour * 24 * 2
	twodays := now.Add(twodaysDiff)
	fmt.Println("Two Days: ", twodays.Format(time.ANSIC))
}
