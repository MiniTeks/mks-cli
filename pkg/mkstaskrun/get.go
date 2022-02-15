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

package mkstaskrun

import (
	"context"
	"fmt"
	"io"

	"github.com/MiniTeks/mks-cli/pkg/mconfig"
	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func getMksTaskRun(mksc *mconfig.Client) *cobra.Command {
	cc := &cobra.Command{
		Use:   "get",
		Short: "Get a MkstaskRun in default namespace",
		RunE: func(cmd *cobra.Command, args []string) error {
			fs, _ := cmd.Flags().GetString("name")
			if fs == "" {
				fmt.Errorf("TaskName not defined")
			}
			obj, err := mksc.Mks.MkscontrollerV1alpha1().MksTaskRuns("default").Get(context.TODO(), fs, metav1.GetOptions{})
			if err != nil {
				return nil
			}
			displayCrt(obj, cmd.OutOrStdout())
			return nil
		},
	}
	cc.Flags().String("name", "", "Name of the taskrun to be deleted")
	return cc
}

func displayCrt(crt *v1alpha1.MksTaskRun, w io.Writer) {
	fmt.Fprintf(w, "name: %s\n", crt.Name)
	fmt.Fprintf(w, "namespace: %s\n", crt.Namespace)
	fmt.Fprintf(w, "taskrunref: %s\n", crt.Spec.TaskRef.Name)
}
