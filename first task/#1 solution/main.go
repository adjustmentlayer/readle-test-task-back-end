package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

type Holiday struct {
	Date      string
	LocalName string
}

func main() {
	monthsUA := []string{
		"Ciчень",
		"Лютий",
		"Березень",
		"Квітень",
		"Травень",
		"Червень",
		"Липень",
		"Серпень",
		"Вересень",
		"Жовтень",
		"Листопад",
		"Грудень",
	}
	var holidays []Holiday
	getJSON("https://date.nager.at/Api/v2/NextPublicHolidays/UA", &holidays)

	var output string

	today := time.Now()
	nextHoliday, _ := time.Parse("2006-01-02", holidays[0].Date)
	nextHolidayName := holidays[0].LocalName

	oneDayEarlier := nextHoliday.AddDate(0, 0, -1)
	oneDayLater := nextHoliday.AddDate(0, 0, 1)

	twoDaysEarlier := nextHoliday.AddDate(0, 0, -2)
	twoDaysLater := nextHoliday.AddDate(0, 0, 2)

	monthNextHoliday := monthsUA[int(nextHoliday.Month())]

	monthOneDayEarlier := monthsUA[int(twoDaysEarlier.Month())]
	monthOneDayLater := monthsUA[int(twoDaysLater.Month())]

	monthTwoDaysEarlier := monthsUA[int(twoDaysEarlier.Month())]
	monthTwoDaysLater := monthsUA[int(twoDaysLater.Month())]

	if today.Equal(nextHoliday) {
		output += "Сьогоднішнє свято " + nextHolidayName + ", " + strconv.Itoa(nextHoliday.Day()) + " " + monthNextHoliday
	} else {
		output += "Наступне свято " + nextHolidayName + ", " + strconv.Itoa(nextHoliday.Day()) + " " + monthNextHoliday
	}
	if nextHoliday.Weekday() == 5 { // Если пятница
		output += ", і вихідні триватимуть 3 дні: "
		output += monthNextHoliday + " " + strconv.Itoa(nextHoliday.Day()) + " - " + monthTwoDaysLater + " " + strconv.Itoa(twoDaysLater.Day())
	}
	if nextHoliday.Weekday() == 1 { // Если понедельник
		output += ", і вихідні триватимуть 3 дні: "
		output += monthTwoDaysEarlier + " " + strconv.Itoa(twoDaysEarlier.Day()) + " - " + monthNextHoliday + " " + strconv.Itoa(nextHoliday.Day())
	}
	if nextHoliday.Weekday() == 6 { // Если суббота
		output += ", і вихідні триватимуть 2 днstrconv.Itoa(і: "
		output += monthNextHoliday + " " + strconv.Itoa(nextHoliday.Day()) + " - " + monthOneDayLater + " " + strconv.Itoa(oneDayLater.Day())
	}
	if nextHoliday.Weekday() == 0 { // Если воскресенье
		output += ", і вихідні триватимуть 2 дні: "
		output += monthOneDayEarlier + " " + strconv.Itoa(oneDayEarlier.Day()) + " - " + monthNextHoliday + " " + strconv.Itoa(nextHoliday.Day())
	}

	fmt.Println(output)

}

func getJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
