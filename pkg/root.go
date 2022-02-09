package cmd

import (
	"os"

	"github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"

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

	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	home = home + "/.kube/config"
	rootCmd.PersistentFlags().StringVar(&CfgFile, "config", home, "k8s config file (default is ${HOME}/.kube/config)")

	cfg, err := clientcmd.BuildConfigFromFlags("", CfgFile)
	cobra.CheckErr(err)
	mksclient, err := versioned.NewForConfig(cfg)
	cobra.CheckErr(err)
	rootCmd.AddCommand(
		mkstask.Command(mksclient),
		mkspipelinerun.Command(mksclient),
		mkstaskrun.Command(mksclient),
	)
}
