package main

import (
	"errors"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func HealthCmd(c *cli.Context) error {
	if c.NArg() < 2 {
		return errors.New("Please add a test case name and a test case file")
		cli.ShowCommandHelp(c, "add")
		return nil
	}

	targetname, outputpath := c.Args().Get(0), c.Args().Get(1)
	protoroot = c.String("protoroot")
	protofile = c.String("protofile")
	protoimport = c.String("protoimport")
	protoimportpath = c.String("protoimportpath")

	if protoroot == "" {
		protoroot = "./"
	}
	_err := os.MkdirAll(protoroot, 0755)
	if _err != nil {
		return _err
	}
	for _, _import := range strings.Split(protoimport, ",") {
		if _import != "" {
			protoimportpath = protoimportpath + ":" + _import
		}
	}
	_, err = os.Stat(outPath)
	if err != nil {
		return errdetails
	}
}

func main() {
	app := cli.NewApp()
	app.Name = "obsidian-client-cli"
	app.Usage = "Obsidian client command line interface"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "Add a new test case",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "protoroot, r",
					Usage: "Root directory for proto files",
				},
				cli.StringFlag{
					Name:  "protofile, p",
					Usage: "Proto file to be used for test case",
				},
				cli.StringFlag{
					Name:  "protoimport, i",
					Usage: "Proto import path",
				},
				cli.StringFlag{
					Name:  "protoimportpath, ip",
					Usage: "Proto import path",
				},
			},
			Action: AddCmd,
		},
		{
			Name:    "Health",
			Aliases: []string{"h"},
			Usage:   "Checking the health of the server",
			Action:  HealthCmd,
		},
	}
	app.Run(os.Args)
}
