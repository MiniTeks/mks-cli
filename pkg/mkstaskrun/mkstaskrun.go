package mkstaskrun

import (
	"github.com/spf13/cobra"
)

// mkstaskrunCmd represents the mkstaskrun command
var mkstaskrunCmd = &cobra.Command{
	Use:   "mkstaskrun",
	Short: "Add create list mkstaskrun",
}

func InitCommand() *cobra.Command {
	mkstaskrunCmd.AddCommand(
		createMksTaskRun(),
		deleteMksTaskRun(),
		listMksTaskRun(),
		getMksTaskRun())
	return mkstaskrunCmd
}
