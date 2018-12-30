package main

import (
	"fmt"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
	uuid "github.com/satori/go.uuid"
)

type NotificationInput struct {
	Message *string
	Date    *graphql.Time
}

type Notification struct {
	// gorm.Model
	ID        uuid.UUID `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	IMessage  string
	IDate     *time.Time
}

func NewNotification(i NotificationInput) Notification {
	id := uuid.Must(uuid.NewV4())
	return Notification{ID: id, IMessage: fmt.Sprint(*i.Message), IDate: &i.Date.Time}
}

func (n Notification) Id() graphql.ID {
	return graphql.ID(n.ID.String())
}

func (n Notification) Message() *string {
	return &n.IMessage
}
func (n Notification) Date() *graphql.Time {
	var t graphql.Time
	if n.IDate != nil {
		t = graphql.Time{Time: *n.IDate}
	}
	return &t
}
