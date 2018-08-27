package model

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/janritter/kvb-api/typedef"

	"github.com/PuerkitoBio/goquery"
)

func GetDeparturesByStationID(id int) []typedef.Departure {

	doc, err := goquery.NewDocument(fmt.Sprintf("https://www.kvb.koeln/generated/?aktion=show&code=%d", id))
	if err != nil {
		log.Fatal(err)
	}

	departures := []typedef.Departure{}

	doc.Find("body > div > table:nth-child(2) > tbody > tr").Each(func(i int, s *goquery.Selection) {
		if i != 0 {
			route := s.Find("td:nth-child(1)").Text()
			direction := s.Find("td:nth-child(2)").Text()
			arrivalTimeString := s.Find("td:nth-child(3)").Text()

			// Build correct time from Sofort and 2 Min
			arrivalTime := -1
			if arrivalTimeString == "Sofort" {
				arrivalTime = 0
			} else {
				arrivalTimeString = strings.Replace(arrivalTimeString, "Min", "", -1)
				arrivalTime, err = strconv.Atoi(strings.TrimSpace(arrivalTimeString))
				if err != nil {
					log.Println("ERROR - Parsing ArrivalTime")
					log.Println(err)
				}
			}

			singleDeparture := typedef.Departure{
				Route:       strings.TrimSpace(route),
				Direction:   direction,
				ArrivalTime: arrivalTime,
			}
			departures = append(departures, singleDeparture)
		}
	})

	return departures
}
