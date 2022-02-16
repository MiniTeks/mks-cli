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

package mkspipelinerun

import (
	"context"
	"fmt"
	"io"

	"github.com/MiniTeks/mks-cli/pkg/mconfig"
	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"

	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func listcommand(mksc *mconfig.Client) *cobra.Command {
	mksPrList := &cobra.Command{
		Use:   "list",
		Short: "List PipelineRuns",
		Annotations: map[string]string{
			"commandType": "main",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			fet, err := mksc.Mks.MkscontrollerV1alpha1().MksPipelineRuns(namespace).List(context.TODO(), v1.ListOptions{})
			if err != nil {
				return nil
			}
			printList(fet, cmd.OutOrStdout())
			return nil
		},
	}
	mksPrList.Flags().StringVar(&namespace, "ns", "default", "NameSpace of MksPipelineRun Resource")

	return mksPrList
}

func printList(fet *v1alpha1.MksPipelineRunList, w io.Writer) {
	fmt.Println("Here is List")
	fmt.Printf("\n")
	for i, obj := range fet.Items {
		fmt.Fprintf(w, "%d %s\n", i+1, obj.GetName())

	}

}
