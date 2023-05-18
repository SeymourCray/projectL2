package usecase

import (
	"dev11/internal/entity"
	"time"
)

type EventsRepository interface {
	AddEvent(e entity.Event) error
	DeleteEvent(id int) error
	UpdateEvent(event entity.Event) error
	GetAllEvents() []entity.Event
}

type EventUseCase struct {
	repository EventsRepository
}

func NewEventUseCase(repository EventsRepository) *EventUseCase {
	return &EventUseCase{repository: repository}
}

func (r *EventUseCase) AddEvent(e entity.Event) error {
	err := r.repository.AddEvent(e)
	if err != nil {
		return err
	}

	return nil
}

func (r *EventUseCase) UpdateEvent(e entity.Event) error {
	err := r.repository.UpdateEvent(e)
	if err != nil {
		return err
	}

	return nil
}

func (r *EventUseCase) DeleteEvent(id int) error {
	err := r.repository.DeleteEvent(id)
	if err != nil {
		return err
	}

	return nil
}

func (r *EventUseCase) FilterEvents(time time.Time) []entity.Event {
	var res []entity.Event
	events := r.repository.GetAllEvents()

	for _, event := range events {
		if time.Before(event.Date.Time) {
			res = append(res, event)
		}
	}

	return res
}

func (r *EventUseCase) name() {

}
