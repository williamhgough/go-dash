package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/williamhgough/go-dash/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
