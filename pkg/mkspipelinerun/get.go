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

func getcommand(mksc *mconfig.Client) *cobra.Command {
	mksPrGet := &cobra.Command{
		Use:   "get",
		Short: "Get PipelineRuns",
		Annotations: map[string]string{
			"commandType": "main",
		},
    
		RunE: func(cmd *cobra.Command, args []string) error {
			get, err := mksc.Mks.MkscontrollerV1alpha1().MksPipelineRuns(namespace).Get(context.TODO(), resourceName, v1.GetOptions{})
			if err != nil {
				return nil
			}
			printPipelineRun(get, cmd.OutOrStdout())
			return nil
		},
	}
	mksPrGet.Flags().StringVar(&resourceName, "rn", "", "Name of MksPipelineRun Resource to be fetched")
	mksPrGet.MarkFlagRequired("rn")
	mksPrGet.Flags().StringVar(&namespace, "ns", "default", "NameSpace of MksPipelineRun Resource")
	return mksPrGet
}

func printPipelineRun(get *v1alpha1.MksPipelineRun, w io.Writer) {
	fmt.Fprintf(w, "name: %s\n", get.Name)
	fmt.Fprintf(w, "namespace: %s\n", get.Namespace)
	fmt.Fprintf(w, "pipelinerunref: %s\n", get.Spec.PipelineRef.Name)
}
