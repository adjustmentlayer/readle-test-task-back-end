package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

type Holiday struct {
	Date      string
	LocalName string
	Name      string
}
type LongWeekend struct {
	StartDate string
	EndDate   string
	DayCount  int
}

func main() {

	var nextPublicHolidays []Holiday
	var longWeekends []LongWeekend
	var publicHolidays []Holiday
	getJSON("https://date.nager.at/Api/v2/NextPublicHolidays/UA", &nextPublicHolidays)
	getJSON("https://date.nager.at/Api/v2/LongWeekend/2020/UA", &longWeekends)
	getJSON("https://date.nager.at/Api/v2/PublicHolidays/2020/UA", &publicHolidays)

	today := time.Now()

	//uncomment the lines below to simulate the onset of the holiday
	/* today, _ := time.Parse("2006-01-02", "2020-08-24") //Independence Day */
	/* today, _ := time.Parse("2006-01-02", "2020-03-08") //International Women`s Day */

	for _, value := range longWeekends {
		startDate, _ := time.Parse("2006-01-02", value.StartDate)
		endDate, _ := time.Parse("2006-01-02", value.EndDate)
		if today.After(startDate.AddDate(0, 0, -1)) && today.Before(endDate.AddDate(0, 0, 1)) {

			name := findLongWeekendName(value.StartDate, value.EndDate, publicHolidays)
			fmt.Println("Today is "+name, ", and the weekend will last", value.DayCount, "days:", startDate.Month(), startDate.Day(), "-", endDate.Month(), endDate.Day())
			os.Exit(0)
		}
	}
	for _, value := range publicHolidays {
		date, _ := time.Parse("2006-01-02", value.Date)
		if today.Equal(date) {
			fmt.Println("Today is "+value.Name+",", date.Month(), date.Day())
			os.Exit(0)
		}
	}

	nextPublicHoliday, _ := time.Parse("2006-01-02", nextPublicHolidays[0].Date)

	for _, value := range longWeekends {
		startDate, _ := time.Parse("2006-01-02", value.StartDate)
		endDate, _ := time.Parse("2006-01-02", value.EndDate)
		if nextPublicHoliday.After(startDate.AddDate(0, 0, -1)) && nextPublicHoliday.Before(endDate.AddDate(0, 0, 1)) {
			name := findLongWeekendName(value.StartDate, value.EndDate, publicHolidays)
			fmt.Println("The next holiday is "+name, ", and the weekend will last", value.DayCount, "days:", startDate.Month(), startDate.Day(), "-", endDate.Month(), endDate.Day())
			os.Exit(0)
		}
	}

	fmt.Println("The next holiday is "+nextPublicHolidays[0].Name+",", nextPublicHoliday.Month(), nextPublicHoliday.Day())

}

func getJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func findLongWeekendName(startDate string, endDate string, publicHolidays []Holiday) string {

	startDate1, _ := time.Parse("2006-01-02", startDate)
	endDate1, _ := time.Parse("2006-01-02", endDate)
	for _, value := range publicHolidays {
		date1, _ := time.Parse("2006-01-02", value.Date)
		if startDate1.Equal(date1) || endDate1.Equal(date1) {
			return value.Name
		}

	}
	return ""
}
