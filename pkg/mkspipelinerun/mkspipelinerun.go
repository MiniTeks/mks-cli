package mkspipelinerun

import (
	"github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
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

func Command(mksclient *versioned.Clientset) *cobra.Command {
	mksPrCmd.AddCommand(
		createcommand(mksclient),
		listcommand(mksclient),
		deletecommand(mksclient),
		getcommand(mksclient),
	)
	return mksPrCmd
}
