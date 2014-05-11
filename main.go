package main

import (
  "github.com/codegangsta/cli"
  "os"
  "hipchat"
)

func main() {
  tken := os.Getenv("TOKEN")
  hc := hipchat.New(tken)

  app := cli.NewApp()
  app.Usage = "Hipchat command line application"
  app.Version = "0.0.1"
  app.Commands = []cli.Command{
    {
      Name:  "room",
      ShortName: "r",
      Usage: "Rooms",
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

            if res != "" {
              println("Message sent")
            }else {
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
