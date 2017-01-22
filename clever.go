package main

import (
	"os"

	"github.com/macrocoders/clever/auth"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "clever"
	app.Usage = "a CLI app for Evernote"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:  "login",
			Usage: "authorize to the Evernote server using OAuth.",
			Action: func(c *cli.Context) {
				auth.Login()
			},
		}}

	app.Run(os.Args)
}
