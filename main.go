package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/99designs/keyring"
	"github.com/urfave/cli/v2"
)

const (
	nameFlag = "name"
)

var (
	kr keyring.Keyring
)

func main() {
	app := &cli.App{
		Name:  "volbo",
		Usage: "simple multi platform vault for storing local secrets",
		Commands: []*cli.Command{
			{
				Name:  "get",
				Usage: "retrieve secret and output to stdout",
				Flags: []cli.Flag{
					NameFlag(),
				},
				Before: OpenKeyRing,
				Action: func(c *cli.Context) error {
					name := c.String(nameFlag)

					v, err := kr.Get(name)
					if err != nil {
						return err
					}

					fmt.Fprintf(os.Stdout, "%s", string(v.Data))

					return nil
				},
			},
			{
				Name:  "set",
				Usage: "store secret received from stdin",
				Flags: []cli.Flag{
					NameFlag(),
				},
				Before: OpenKeyRing,
				Action: func(c *cli.Context) error {
					name := c.String(nameFlag)
					reader := bufio.NewReader(os.Stdin)
					data, err := reader.ReadString('\n')
					if err != nil {
						return err
					}

					if err := kr.Set(keyring.Item{
						Key:  name,
						Data: []byte(data),
					}); err != nil {
						return err
					}

					return nil
				},
			},
			{
				Name:  "del",
				Usage: "remove secret from vault",
				Flags: []cli.Flag{
					NameFlag(),
				},
				Before: OpenKeyRing,
				Action: func(c *cli.Context) error {
					name := c.String(nameFlag)

					if err := kr.Remove(name); err != nil {
						return err
					}

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}

func OpenKeyRing(c *cli.Context) (err error) {
	kr, err = keyring.Open(keyring.Config{
		ServiceName: "volbo",
	})
	return
}

func NameFlag() cli.Flag {
	return &cli.StringFlag{
		Name:     nameFlag,
		Required: true,
	}
}
