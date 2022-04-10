package cc

import (
	"io"
	"os"
)

type CommonCtx struct {
	outWriter io.Writer
	errWriter io.Writer
}

func NewCommonCtx() *CommonCtx {
	return &CommonCtx{
		outWriter: os.Stdout,
		errWriter: os.Stderr,
	}
}

func (c *CommonCtx) OutWriter() io.Writer {
	return c.outWriter
}

func (c *CommonCtx) ErrWriter() io.Writer {
	return c.errWriter
}
