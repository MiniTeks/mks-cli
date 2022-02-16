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

	"github.com/MiniTeks/mks-cli/pkg/mconfig"
	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"

	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog/v2"
)

func createcommand(mksc *mconfig.Client) *cobra.Command {
	mksPrCreate := &cobra.Command{
		Use:   "create",
		Short: "Create PipelineRuns",
		Annotations: map[string]string{
			"commandType": "main",
		},
		Run: func(cmd *cobra.Command, args []string) {
			deployment := &v1alpha1.MksPipelineRun{
				TypeMeta:   v1.TypeMeta{Kind: "MksPipelineRun"},
				ObjectMeta: v1.ObjectMeta{Name: resourceName},
				Spec:       v1alpha1.MksPipelineRunSpec{PipelineRef: v1alpha1.MksPipelineRunRef{Name: pipelineRunRef}},
			}

			crt, err := mksc.Mks.MkscontrollerV1alpha1().MksPipelineRuns(namespace).Create(context.TODO(), deployment, v1.CreateOptions{})
			if err != nil {
				klog.Fatalf("Create MksPipelineRun failed!", err.Error())
			}
			fmt.Fprintf(cmd.OutOrStdout(), crt.Name)
		},
	}
	mksPrCreate.Flags().StringVar(&resourceName, "n", "", "Name of MksPipelineRun Resource")
	mksPrCreate.MarkFlagRequired("n")
	mksPrCreate.Flags().StringVar(&pipelineRunRef, "pr", "", "Name for PipelineRun Ref")
	mksPrCreate.MarkFlagRequired("pr")
	mksPrCreate.Flags().StringVar(&namespace, "ns", "default", "NameSpace of MksPipelineRun Resource")
	return mksPrCreate
}
