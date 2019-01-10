package main

import (
	graphql "github.com/graph-gophers/graphql-go"
	uuid "github.com/satori/go.uuid"
)

// NewNotification ...
func NewNotification(i notificationInput) Notification {
	id := uuid.Must(uuid.NewV4())
	seen := false
	return Notification{
		IID:          id,
		IMessage:     i.Message,
		IPrincipal:   i.Principal,
		IChannel:     i.Channel,
		IReference:   i.Reference,
		IReferenceID: i.ReferenceID,
		IDate:        i.Date.Time,
		ISeen:        &seen,
	}
}

// ID ...
func (n Notification) ID() graphql.ID {
	return graphql.ID(n.IID.String())
}

// Message ...
func (n Notification) Message() string {
	return n.IMessage
}

// Principal ...
func (n Notification) Principal() *string {
	return n.IPrincipal
}

// Channel ...
func (n Notification) Channel() *string {
	return n.IChannel
}

// Reference ...
func (n Notification) Reference() *string {
	return n.IReference
}

// ReferenceID ...
func (n Notification) ReferenceID() *string {
	return n.IReferenceID
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
