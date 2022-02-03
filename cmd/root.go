/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/MiniTeks/mks-cli/pkg/mkstask"
	"os"

	"github.com/MiniTeks/mks-cli/pkg/mkspipelinerun"
	"github.com/spf13/cobra"
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
	)
	home, err := os.UserHomeDir()
	cobra.CheckErr(err)
	home = home + "/.kube/config"
	rootCmd.PersistentFlags().StringVar(&CfgFile, "config", home, "k8s config file (default is ${HOME}/.kube/config)")
}
