// SPDX-License-Identifier: Apache-2.0
// SPDX-FileCopyrightText: 2022 Satyam Bhardwaj <sabhardw@redhat.com>
// SPDX-FileCopyrightText: 2022 Utkarsh Chaurasia <uchauras@redhat.com>
// SPDX-FileCopyrightText: 2022 Avinal Kumar <avinkuma@redhat.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//    http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"os"

	"github.com/MiniTeks/mks-cli/pkg/mkspipelinerun"
	"github.com/MiniTeks/mks-cli/pkg/mkstask"
	"github.com/MiniTeks/mks-cli/pkg/mkstaskrun"
	"github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
	"k8s.io/client-go/tools/clientcmd"
)

var CfgFile, home string

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

func GenerateDocs() {
	err := doc.GenMarkdownTree(rootCmd, "/tmp")
	cobra.CheckErr(err)
}
