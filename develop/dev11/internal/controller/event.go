package controller

import (
	"dev11/internal/entity"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
Реализовать middleware для логирования запросов
*/

type EventUseCase interface {
	FilterEvents(time time.Time) []entity.Event
	AddEvent(e entity.Event) error
	DeleteEvent(id int) error
	UpdateEvent(event entity.Event) error
}

type EventController struct {
	eventUseCase EventUseCase
}

type eventResponse struct {
	Result []entity.Event `json:"result"`
}

func NewEventController(e EventUseCase) *EventController {
	return &EventController{e}
}

func (c *EventController) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	var event entity.Event

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		log.Println(err)
		http.Error(w, "wrong format", http.StatusBadRequest)
		return
	}

	err := c.eventUseCase.AddEvent(event)
	if err != nil {
		http.Error(w, "usecase error", http.StatusServiceUnavailable)
		return
	}

	response, err := json.Marshal(event)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = fmt.Fprint(w, string(response))
	if err != nil {
		log.Println(err)
	}
}

func (c *EventController) UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	var event entity.Event

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "wrong format", http.StatusBadRequest)
		return
	}

	err := c.eventUseCase.UpdateEvent(event)
	if err != nil {
		http.Error(w, "usecase error", http.StatusServiceUnavailable)
		return
	}

	response, err := json.Marshal(event)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = fmt.Fprint(w, string(response))
	if err != nil {
		log.Println(err)
	}
}

func (c *EventController) DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	var event entity.Event

	if err := json.NewDecoder(r.Body).Decode(&event); err != nil {
		http.Error(w, "wrong format", http.StatusBadRequest)
		return
	}

	err := c.eventUseCase.DeleteEvent(event.UserID)
	if err != nil {
		http.Error(w, "usecase error", http.StatusServiceUnavailable)
		return
	}

	response, err := json.Marshal(event)
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = fmt.Fprint(w, string(response))
	if err != nil {
		log.Println(err)
	}
}

func (c *EventController) EventForDayHandler(w http.ResponseWriter, r *http.Request) {
	var events []entity.Event

	events = c.eventUseCase.FilterEvents(time.Now().AddDate(0, 0, -1))

	response, err := json.Marshal(eventResponse{events})
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = fmt.Fprint(w, string(response))
	if err != nil {
		log.Println(err)
	}
}

func (c *EventController) EventForWeekHandler(w http.ResponseWriter, r *http.Request) {
	var events []entity.Event

	events = c.eventUseCase.FilterEvents(time.Now().AddDate(0, 0, -7))

	response, err := json.Marshal(eventResponse{events})
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = fmt.Fprint(w, string(response))
	if err != nil {
		log.Println(err)
	}
}

func (c *EventController) EventForMonthHandler(w http.ResponseWriter, r *http.Request) {
	var events []entity.Event

	events = c.eventUseCase.FilterEvents(time.Now().AddDate(0, -1, 0))

	response, err := json.Marshal(eventResponse{events})
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = fmt.Fprint(w, string(response))
	if err != nil {
		log.Println(err)
	}
}

func (c *EventController) RegisterHandlers() {
	http.HandleFunc("/create_event", c.CreateEventHandler)
	http.HandleFunc("/update_event", c.UpdateEventHandler)
	http.HandleFunc("/delete_event", c.DeleteEventHandler)
	http.HandleFunc("/events_for_day", c.EventForDayHandler)
	http.HandleFunc("/events_for_week", c.EventForWeekHandler)
	http.HandleFunc("/events_for_month", c.EventForMonthHandler)
}
