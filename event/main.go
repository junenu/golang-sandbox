package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Event struct {
	Name    string    `json:"name"`
	Date    time.Time `json:"date"`
	details string    `json:"details"`
}

var events []Event

func (e Event) MarshalJSON() ([]byte, error) {
	type Alias Event
	return json.Marshal(&struct {
		Date string `json:"date"`
		*Alias
	}{
		Date:  e.Date.Format("2006/01/02"),
		Alias: (*Alias)(&e),
	})
}

func (e *Event) UnmarshalJSON(data []byte) error {
	type Alias Event
	aux := &struct {
		Date string `json:"date"`
		*Alias
	}{
		Alias: (*Alias)(e),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	date, err := time.Parse("2006/01/02", aux.Date)
	if err != nil {
		return err
	}
	e.Date = date
	return nil
}

func registerEvent(c *gin.Context) {
	var newEvent Event // Event構造体をグローバルに定義したので、ここでは型宣言のみ
	if err := c.ShouldBindJSON(&newEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	events = append(events, newEvent)
	c.JSON(http.StatusCreated, newEvent)
}

func getEvents(c *gin.Context) {
	now := time.Now()
	var validEvents []Event
	for _, event := range events {
		if event.Date.After(now.AddDate(0, 0, -7)) {
			validEvents = append(validEvents, event)
		}
	}
	c.JSON(http.StatusOK, validEvents)
}

func deletePastEvents() {
	for {
		now := time.Now()
		events = filterEvents(events, func(event Event) bool {
			return event.Date.After(now)
		})
		time.Sleep(24 * time.Hour) // 24時間ごとに実行
	}
}

func filterEvents(events []Event, predicate func(Event) bool) []Event {
	var filteredEvents []Event
	for _, event := range events {
		if predicate(event) {
			filteredEvents = append(filteredEvents, event)
		}
	}
	return filteredEvents
}

func main() {
	router := gin.Default()
	router.POST("/events", registerEvent)
	router.GET("/events", getEvents)
	go deletePastEvents() // 過去のイベント削除をバックグラウンドで実行
	router.Run(":8080")
}

/// ❯ curl -X POST -H "Content-Type: application/json" -d '{
//  "name": "ミャーの勉強会",
//  "date": "2024/12/31",
//  "details": "ミャー"
//}' http://localhost:8080/events
// {"date":"2024/12/31","name":"ミャーの勉強会"}

// {"name":"ミャーの勉強会","date":"2024/12/31","details":"ミャー"}
