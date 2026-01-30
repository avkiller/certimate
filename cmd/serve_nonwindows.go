//go:build !windows
// +build !windows

package cmd

import (
	"github.com/pocketbase/pocketbase"
)

func Serve(app *pocketbase.PocketBase) error {
	return app.Start()
}
