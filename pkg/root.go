package cmd

import (
	"os"

	"github.com/spf13/cobra"

	"github.com/MiniTeks/mks-cli/pkg/mkspipelinerun"
	"github.com/MiniTeks/mks-cli/pkg/mkstask"
	"github.com/MiniTeks/mks-cli/pkg/mkstaskrun"
)

var CfgFile string

var rootCmd = &cobra.Command{
	Use:   "mks",
	Short: "mks is a cli client to interact with mks-server ",
	Long:  "mks is a cli client to interact with mks-server ",
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {

	rootCmd.AddCommand(
		mkstask.Command(),
		mkspipelinerun.Command(CfgFile),
		mkstaskrun.InitCommand(),
	)
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	home = home + "/.kube/config"
	rootCmd.PersistentFlags().StringVar(&CfgFile, "config", home, "k8s config file (default is ${HOME}/.kube/config)")
}
