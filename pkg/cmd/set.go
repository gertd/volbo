package cmd

import (
	"bytes"
	"os"

	"github.com/gertd/volbo/pkg/cc"
	"github.com/gertd/volbo/pkg/keyring"
)

type SetCmd struct {
	Key   string   `arg:"" name:"key" required:"" help:"secret name"`
	Value string   `xor:"group" required:"" name:"value" help:"secret value"`
	Input *os.File `xor:"group" required:"" name:"file" type:"existingfile" help:"secret value from file or stdin"`
}

func (cmd *SetCmd) Run(c *cc.CommonCtx) error {
	var (
		token string
	)

	if cmd.Input != nil {
		buf := new(bytes.Buffer)
		if _, err := buf.ReadFrom(cmd.Input); err != nil {
			return err
		}
		token = buf.String()
	} else if cmd.Value != "" {
		token = cmd.Value
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
