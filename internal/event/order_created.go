package event

import (
	"time"

	"github.com/devfullcycle/20-CleanArch/pkg/events"
)

type OrderCreated struct {
	Name     string
	Payload  interface{}
	DateTime time.Time
}

func NewOrderCreated() *OrderCreated {
	return &OrderCreated{
		Name:     "OrderCreated",
		DateTime: time.Now(),
	}
}

func (e *OrderCreated) GetName() string {
	return e.Name
}

func (e *OrderCreated) GetPayload() interface{} {
	return e.Payload
}

func (e *OrderCreated) SetPayload(payload interface{}) {
	e.Payload = payload
}

func (e *OrderCreated) GetDateTime() time.Time {
	return e.DateTime
}

var _ events.EventInterface = (*OrderCreated)(nil)
