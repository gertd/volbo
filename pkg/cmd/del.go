package cmd

import (
	"github.com/gertd/volbo/pkg/cc"
	"github.com/gertd/volbo/pkg/keyring"
)

type DelCmd struct {
	Key string `arg:"" name:"key" required:"" help:"secret name"`
}

func (cmd *DelCmd) Run(c *cc.CommonCtx) error {
	kr, err := keyring.New(cmd.Key)
	if err != nil {
		return err
	}

	if err := kr.DelToken(); err != nil {
		return err
	}

	return nil
}
