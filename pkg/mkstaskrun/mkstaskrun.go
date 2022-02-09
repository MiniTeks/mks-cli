package mkstaskrun

import (
	"github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	"github.com/spf13/cobra"
)

// mkstaskrunCmd represents the mkstaskrun command
var mkstaskrunCmd = &cobra.Command{
	Use:   "mkstaskrun",
	Short: "Add create list mkstaskrun",
}

func Command(mksclient *versioned.Clientset) *cobra.Command {
	mkstaskrunCmd.AddCommand(
		createMksTaskRun(mksclient),
		deleteMksTaskRun(mksclient),
		listMksTaskRun(mksclient),
		getMksTaskRun(mksclient))
	return mkstaskrunCmd
}
