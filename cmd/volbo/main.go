package main

import (
	"github.com/alecthomas/kong"
	"github.com/gertd/volbo/pkg/cc"
	"github.com/gertd/volbo/pkg/cmd"
	"github.com/gertd/volbo/pkg/x"
)

func main() {
	cli := cmd.CLI{}
	kongCtx := kong.Parse(&cli,
		kong.Name(x.AppName),
		kong.Description(x.AppDescription),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			NoAppSummary:        false,
			Summary:             false,
			Compact:             true,
			Tree:                false,
			FlagsLast:           true,
			Indenter:            kong.SpaceIndenter,
			NoExpandSubcommands: false,
		}),
	)

	ctx := cc.NewCommonCtx()

	if err := kongCtx.Run(ctx); err != nil {
		kongCtx.FatalIfErrorf(err)
	}
}
