package main

import (
	"fmt"

	graphql "github.com/graph-gophers/graphql-go"
	uuid "github.com/satori/go.uuid"
)

// NewNotification ...
func NewNotification(i notificationInput) Notification {
	id := uuid.Must(uuid.NewV4())
	seen := false
	return Notification{
		IID:        id,
		IMessage:   fmt.Sprint(*i.Message),
		IPrincipal: i.Principal,
		IChannel:   i.Channel,
		IReference: i.Reference,
		IDate:      i.Date.Time,
		ISeen:      &seen,
	}
}

// ID ...
func (n Notification) ID() graphql.ID {
	return graphql.ID(n.IID.String())
}

// Message ...
func (n Notification) Message() *string {
	return &n.IMessage
}

// Principal ...
func (n Notification) Principal() graphql.ID {
	return graphql.ID(n.IPrincipal)
}

// Channel ...
func (n Notification) Channel() graphql.ID {
	return graphql.ID(n.IChannel)
}

// Reference ...
func (n Notification) Reference() graphql.ID {
	return graphql.ID(n.IReference)
}

// Date ...
func (n Notification) Date() graphql.Time {
	return graphql.Time{Time: n.IDate}
}

// Seen ...
func (n Notification) Seen() bool {
	if n.ISeen != nil {
		return *n.ISeen
	}
	return false
}
