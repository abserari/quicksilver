package main

import (
	"github.com/jinzhu/gorm"
	"github.com/qor/transition"
)

type Order struct {
	ID uint
	transition.Transition
}

func main() {

	var OrderStateMachine = transition.New(&Order{})

	OrderStateMachine.Initial("draft")

	OrderStateMachine.State("checkout")

	OrderStateMachine.State("paid").Enter(func(order interface{}, tx *gorm.DB) error {
		return nil
	}).Exit(func(order interface{}, tx *gorm.DB) error {
		return nil
	})

	OrderStateMachine.State("cancelled")
	OrderStateMachine.State("paid_cancelled")

	OrderStateMachine.Event("checkout").To("checkout").From("draft")

	OrderStateMachine.Event("paid").To("paid").From("checkout").Before(func(order interface{}, tx *gorm.DB) error {
		return nil
	}).After(func(order interface{}, tx *gorm.DB) error {
		return nil
	})

	// Different state transitions for one event
	cancellEvent := OrderStateMachine.Event("cancel")
	cancellEvent.To("cancelled").From("draft", "checkout")
	cancellEvent.To("paid_cancelled").From("paid").After(func(order interface{}, tx *gorm.DB) error {
		// Refund
		return nil
	})
}
