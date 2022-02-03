package mkstask

import (
	"github.com/spf13/cobra"
)

var mksTaskCmd = &cobra.Command{
	Use:   "mkstask",
	Short: "mkstask <option>",
	Long:  "mkstask is to be used to create, get, delete, update, list mksTask resources",
}

func Command() *cobra.Command {
	return mksTaskCmd

}

func init() {
	Command().AddCommand(MksTaskCreate())
	Command().AddCommand(MksTaskGet())
	Command().AddCommand(MksTaskList())
	Command().AddCommand(MksTaskDelete())
	Command().AddCommand(MksTaskUpdate())
}
