package main

import (
	"fmt"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/urfave/cli"
)

type query struct {
	db *DB
}

func (q *query) Notifications() []Notification {
	var notifications []Notification
	q.db.db.Find(&notifications)
	return notifications
}
func (q *query) Notification(params struct{ ID graphql.ID }) *Notification {
	var n Notification
	q.db.db.First(&n, params.ID)
	return &n
}
func (q *query) CreateNotification(params struct{ Input NotificationInput }) Notification {
	n := NewNotification(params.Input)
	q.db.db.Create(&n)
	return n
}

func ServerCommand() cli.Command {
	return cli.Command{
		Name: "server",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "db",
				Usage:  "Connection url to database (GORM)",
				EnvVar: "DATABASE_URL",
			},
			cli.StringFlag{
				Name:   "p,port",
				Usage:  "Server port to bind",
				EnvVar: "PORT",
				Value:  "80",
			},
		},
		Action: func(c *cli.Context) error {
			databaseUrl := c.String("db")
			port := c.String("port")

			if err := startServer(databaseUrl, port); err != nil {
				return cli.NewExitError(err, 1)
			}

			return nil
		},
	}
}
func startServer(urlString, port string) error {
	s := `
		scalar Time
		schema {
			query: Query
			mutation: Mutation
		}
		type EventStoreNotification {
			id: ID!
			message: String
			date: Time
		}
		type Query {
			notifications: [EventStoreNotification!]!
			notification(id: ID!): EventStoreNotification
		}
		input EventStoreNotificationInput {
			message: String
			date: Time
		}
		type Mutation {
			createNotification(input: EventStoreNotificationInput!): EventStoreNotification!
		}
	`

	db := NewDBWithString(urlString)
	defer db.Close()
	db.AutoMigrate(&Notification{})

	schema := graphql.MustParseSchema(s, &query{db})
	http.Handle("/graphql", &relay.Handler{Schema: schema})

	fmt.Println("starting on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	return nil
}
