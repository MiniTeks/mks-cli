package mkstask

import (
	"github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	"github.com/spf13/cobra"
)

var mksTaskCmd = &cobra.Command{
	Use:   "mkstask",
	Short: "mkstask <option>",
	Long:  "mkstask is to be used to create, get, delete, update, list mksTask resources",
}

func Command(mksclient *versioned.Clientset) *cobra.Command {
	mksTaskCmd.AddCommand(
		MksTaskCreate(mksclient),
		MksTaskGet(mksclient),
		MksTaskList(mksclient),
		MksTaskDelete(mksclient),
		MksTaskUpdate(mksclient))
	return mksTaskCmd
}
