package cmd

import (
	"fmt"

	"github.com/gertd/volbo/pkg/cc"
	"github.com/gertd/volbo/pkg/keyring"
)

type GetCmd struct {
	Key string `arg:"" name:"key" required:"" help:"secret name"`
}

func (cmd *GetCmd) Run(c *cc.CommonCtx) error {
	kr, err := keyring.New(cmd.Key)
	if err != nil {
		return err
	}

	token, err := kr.GetToken()
	if err != nil {
		return err
	}

	fmt.Fprintln(c.OutWriter(), token)

	return nil
}
