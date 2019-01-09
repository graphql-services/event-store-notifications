package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/urfave/cli"
)

// ServerCommand ...
func ServerCommand() cli.Command {
	return cli.Command{
		Name: "server",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "db",
				Usage:  "Connection url to database (GORM)",
				EnvVar: "DATABASE_URL",
				Value:  "sqlite3://test.db",
			},
			cli.StringFlag{
				Name:   "p,port",
				Usage:  "Server port to bind",
				EnvVar: "PORT",
				Value:  "80",
			},
		},
		Action: func(c *cli.Context) error {
			databaseURL := c.String("db")

			if databaseURL == "" {
				return cli.NewExitError(fmt.Errorf("database url must be provided"), 1)
			}

			port := c.String("port")

			if err := startServer(databaseURL, port); err != nil {
				return cli.NewExitError(err, 1)
			}

			return nil
		},
	}
}

func startServer(urlString, port string) error {
	dat, err := ioutil.ReadFile("schema.graphql")
	if err != nil {
		return err
	}
	s := string(dat)

	db := NewDBWithString(urlString)
	defer db.Close()
	db.AutoMigrate(&Notification{})

	schema := graphql.MustParseSchema(s, &query{db})
	http.Handle("/graphql", &relay.Handler{Schema: schema})

	fmt.Println("starting on port: " + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

	return nil
}
