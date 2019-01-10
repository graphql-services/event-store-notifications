package main

import (
	"fmt"
	"time"

	graphql "github.com/graph-gophers/graphql-go"
)

type query struct {
	db *DB
}

type notificationsParams struct {
	Principal, Channel, Reference *string
	Seen                          *bool
	Limit                         *int32
}

// Notifications ...
func (q *query) Notifications(params notificationsParams) ([]Notification, error) {
	var notifications []Notification
	query := q.db.db
	if params.Principal != nil {
		query = query.Where(&Notification{IPrincipal: *params.Principal})
	}
	if params.Channel != nil {
		query = query.Where(&Notification{IChannel: *params.Channel})
	}
	if params.Reference != nil {
		query = query.Where(&Notification{IReference: *params.Reference})
	}
	if params.Seen != nil {
		fmt.Println("seen!!")
		query = query.Where(&Notification{ISeen: params.Seen})
	}
	if params.Limit != nil {
		query = query.Limit(*params.Limit)
	}
	query = query.Order("IDate desc")
	query.Find(&notifications)
	return notifications, query.Error
}

// Notification ...
func (q *query) Notification(params struct{ ID graphql.ID }) *Notification {
	var n Notification
	if q.db.db.Where("i_id = ?", params.ID).First(&n).RecordNotFound() {
		return nil
	}
	return &n
}

type notificationInput struct {
	Message   *string
	Principal string
	Channel   string
	Reference string
	Date      graphql.Time
}

// CreateNotification ...
func (q *query) CreateNotification(params struct{ Input notificationInput }) Notification {
	n := NewNotification(params.Input)
	q.db.db.Create(&n)
	fmt.Println(n)
	return n
}

// SeenNotification ...
func (q *query) SeenNotification(params struct{ ID graphql.ID }) *Notification {
	var n Notification
	if q.db.db.Where("i_id =?", params.ID).First(&n).RecordNotFound() {
		return nil
	}
	// seen := true
	// n.ISeen = &seen
	q.db.db.Model(&n).Where(&Notification{IID: n.IID}).Update("ISeen", true)
	return &n
}

// SeenNotificationsParams ...
type SeenNotificationsParams struct {
	Principal graphql.ID
	Channel   *graphql.ID
	Reference *graphql.ID
	Date      time.Time
}

// SeenNotifications ...
func (q *query) SeenNotifications(params SeenNotificationsParams) bool {
	// n := NewNotification(params.)
	// q.db.db.Create(&n)
	// return n
	return true
}
