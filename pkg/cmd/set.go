package cmd

import (
	"bytes"
	"log"
	"os"

	"github.com/gertd/volbo/pkg/cc"
	"github.com/gertd/volbo/pkg/keyring"
	"github.com/pkg/errors"
	"golang.org/x/term"
)

type SetCmd struct {
	Key    string `arg:"" name:"key" required:"" help:"secret name"`
	Prompt bool   `xor:"group" required:"" name:"prompt" help:"secret value read from secure prompt"`
	Stdin  bool   `xor:"group" required:"" name:"stdin" help:"secret value read from stdin"`
	File   string `xor:"group" required:"" name:"file" type:"existingfile" help:"secret value read from file"`
}

func (cmd *SetCmd) Run(c *cc.CommonCtx) error {
	var (
		token string
	)

	switch {
	case cmd.Stdin:
		buf := new(bytes.Buffer)
		if _, err := buf.ReadFrom(os.Stdin); err != nil {
			return errors.Wrapf(err, "read from stdin")
		}
		token = buf.String()

	case cmd.File != "":
		buf, err := os.ReadFile(cmd.File)
		if err != nil {
			return errors.Wrapf(err, "read file [%s]", cmd.File)
		}
		token = string(buf)

	case cmd.Prompt:
		buf, err := term.ReadPassword(0)
		if err != nil {
			log.Fatalln(err)
		}
		token = string(buf)
	}

	kr, err := keyring.New(cmd.Key)
	if err != nil {
		return err
	}

	if err := kr.SetToken(token); err != nil {
		return err
	}

	return nil
}
