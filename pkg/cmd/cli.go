package cmd

type CLI struct {
	Get     GetCmd     `cmd:"" help:"get secret"`
	Set     SetCmd     `cmd:"" help:"set secret"`
	Del     DelCmd     `cmd:"" help:"delete secret"`
	Version VersionCmd `cmd:"" help:"version info"`
}
