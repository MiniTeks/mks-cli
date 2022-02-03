package mkspipelinerun

import (
	"github.com/spf13/cobra"
)

//flags var
var resourceName string
var pipelineRunRef string
var namespace string
var cfgFile string

// mksPrCmd represents the mkspipelinerun command

var mksPrCmd = &cobra.Command{
	Use:   "mkspipelinerun",
	Short: "Manage PipelineRuns",
}

func Command() *cobra.Command {
	return mksPrCmd
}

func init() {
	Command().AddCommand(
		createcommand(),
		listcommand(),
		deletecommand(),
		getcommand(),
	)

}
