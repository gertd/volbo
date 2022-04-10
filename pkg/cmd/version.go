package cmd

import (
	"fmt"

	"github.com/gertd/volbo/pkg/cc"
	"github.com/gertd/volbo/pkg/version"
	"github.com/gertd/volbo/pkg/x"
)

type VersionCmd struct{}

func (cmd *VersionCmd) Run(c *cc.CommonCtx) error {
	fmt.Fprintf(c.OutWriter(), "%s - %s (%s)\n",
		x.AppName,
		version.GetInfo().String(),
		x.AppVersionTag,
	)
	return nil
}
