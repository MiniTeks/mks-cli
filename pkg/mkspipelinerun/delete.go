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
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func deletecommand(mksc *mconfig.Client) *cobra.Command {
	mksPrDelete := &cobra.Command{
		Use:   "delete",
		Short: "Delete PipelineRuns",
		Annotations: map[string]string{
			"commandType": "main",
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			delErr := mksc.Mks.MkscontrollerV1alpha1().MksPipelineRuns(namespace).Delete(context.TODO(), resourceName, v1.DeleteOptions{})
			if delErr != nil {
				return nil
			}
			fmt.Println("Mks PipelineRun ", resourceName, " deleted")
			return nil
		},
	}
	mksPrDelete.Flags().StringVar(&resourceName, "rn", "", "Name of MksPipelineRun Resource to be deleted")
	mksPrDelete.MarkFlagRequired("rn")
	mksPrDelete.Flags().StringVar(&namespace, "ns", "default", "NameSpace of MksPipelineRun Resource")

	return mksPrDelete
}
