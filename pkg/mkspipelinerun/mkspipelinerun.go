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

func Command(cfgPath string) *cobra.Command {
	cfgFile = cfgPath
	return mksPrCmd
}

func init() {
	Command(cfgFile).AddCommand(
		createcommand(),
		listcommand(),
		deletecommand(),
		getcommand(),
	)

}
