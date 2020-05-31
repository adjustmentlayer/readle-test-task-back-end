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

	var output string
	var holidays []Holiday
	getJSON("https://date.nager.at/Api/v2/NextPublicHolidays/UA", &holidays)

	relativeOffset := []int{-1, -2, 0, 0, 0, 2, 1}
	cases := []string{"2 доби", "3 доби", "1 день", "1 день", "1 день", "3 доби", "2 доби"}
	monthsUA := []string{"Ciчень", "Лютий", "Березень", "Квітень", "Травень", "Червень", "Липень", "Серпень", "Вересень", "Жовтень", "Листопад", "Грудень"}

	today := time.Now()
	nextHoliday, _ := time.Parse("2006-01-02", holidays[0].Date)
	nextHolidayName := holidays[0].LocalName

	if today.Equal(nextHoliday) {
		output += "Сьогоднішнє свято "
	} else {
		output += "Наступне свято "
	}
	output += nextHolidayName + " "
	output += ", і вихідні триватимуть " + cases[int(nextHoliday.Weekday())] + ": "

	if relativeOffset[int(nextHoliday.Weekday())] < 0 {
		start := nextHoliday.AddDate(0, 0, relativeOffset[int64(nextHoliday.Weekday())])
		end := nextHoliday
		output += monthsUA[start.Month()] + " " + strconv.Itoa(start.Day()) + " - " + monthsUA[end.Month()] + " " + strconv.Itoa(end.Day())

	} else {
		start := nextHoliday
		end := nextHoliday.AddDate(0, 0, relativeOffset[int64(nextHoliday.Weekday())])
		output += monthsUA[start.Month()] + " " + strconv.Itoa(start.Day()) + " - " + monthsUA[end.Month()] + " " + strconv.Itoa(end.Day())
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
