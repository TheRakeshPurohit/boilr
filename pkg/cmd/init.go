package cmd

import (
	"errors"
	"fmt"

	cli "github.com/spf13/cobra"
	"github.com/tmrts/tmplt/pkg/tmplt"
	"github.com/tmrts/tmplt/pkg/util/exit"
	"github.com/tmrts/tmplt/pkg/util/osutil"
)

var (
	ErrUninitializedTmpltDir = errors.New("tmplt: .tmplt directory is not initialized")
)

// TODO remove?
var Init = &cli.Command{
	Use:   "init",
	Short: "Initialize directories required by tmplt (By default done by installation script)",
	Run: func(c *cli.Command, _ []string) {
		// Check if .config/tmplt exists
		if exists, err := osutil.DirExists(tmplt.Configuration.TemplateDirPath); exists {
			if shouldRecreate := GetBoolFlag(c, "force"); !shouldRecreate {
				exit.Error(ErrUninitializedTmpltDir)
			}
		} else if err != nil {
			exit.Error(fmt.Errorf("init: %s", err))
		}

		if err := osutil.CreateDirs(tmplt.Configuration.TemplateDirPath); err != nil {
			exit.Error(err)
		}

		exit.OK("Initialization complete")
	},
}
