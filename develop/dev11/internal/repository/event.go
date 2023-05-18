package repository

import (
	"dev11/internal/entity"
	"errors"
	"sync"
)

type CashedEvents struct {
	sync.RWMutex
	events map[int]entity.Event
}

func NewCashedEvents() *CashedEvents {
	return &CashedEvents{events: make(map[int]entity.Event)}
}

func (r *CashedEvents) AddEvent(e entity.Event) error {
	r.Lock()

	defer r.Unlock()

	if _, ok := r.events[e.UserID]; ok {
		return errors.New("event with such id already exists")
	}

	r.events[e.UserID] = e

	return nil
}

func (r *CashedEvents) DeleteEvent(id int) error {
	r.Lock()

	defer r.Unlock()

	delete(r.events, id)

	return nil
}

func (r *CashedEvents) UpdateEvent(e entity.Event) error {
	r.Lock()

	defer r.Unlock()

	e, ok := r.events[e.UserID]
	if !ok {
		return errors.New("no event with such id")
	}

	r.events[e.UserID] = e

	return nil
}

func (r *CashedEvents) GetAllEvents() []entity.Event {
	res := make([]entity.Event, len(r.events))

	r.RLock()

	defer r.RUnlock()

	for _, event := range r.events {
		res = append(res, event)
	}

	return res
}
