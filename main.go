package main

import (
	"github.com/codegangsta/cli"
	"hipchat"
	"os"
)

func main() {
	tken := os.Getenv("TOKEN")
	hc := hipchat.New(tken)

	app := cli.NewApp()
	app.Usage = "Hipchat command line application"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:      "user",
			ShortName: "u",
			Usage:     "Users",
			Subcommands: []cli.Command{
				{
					Name:  "message",
					Usage: "send a message to a user",
					Action: func(c *cli.Context) {
						if len(c.Args()) < 2 {
							println("user name and message are required.")
							os.Exit(1)
						}

						user := c.Args().First()
						msg := c.Args()[1]
						from := os.Getenv("USER")
						res, err := hc.MessageUser(user, msg, from)
						if err != nil {
							println(err)
							os.Exit(1)
						}

						if res != "204" {
							println("Message sent")
						} else {
							println("RoomMessageError: ", res)
						}
					},
				},
			},
		},
		{
			Name:      "room",
			ShortName: "r",
			Usage:     "Rooms",
			Subcommands: []cli.Command{
				{
					Name:  "message",
					Usage: "send a message to a room",
					Action: func(c *cli.Context) {
						if len(c.Args()) < 2 {
							println("room name and message are required.")
							os.Exit(1)
						}

						room := c.Args().First()
						msg := c.Args()[1]
						from := os.Getenv("USER")
						res, err := hc.MessageRoom(room, msg, from)
						if err != nil {
							println(err)
							os.Exit(1)
						}

						if res != "204" {
							println("Message sent")
						} else {
							println("RoomMessageError: ", res)
						}
					},
				},
				{
					Name:  "list",
					Usage: "List rooms",
					Action: func(c *cli.Context) {
						println("list rooms: ", c.Args().First())
					},
				},
			},
		},
	}
	app.Run(os.Args)
}
