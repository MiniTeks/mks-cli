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

	"github.com/MiniTeks/mks-cli/pkg/mconfig"
	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	"github.com/spf13/cobra"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func createMksTaskRun(mksc *mconfig.Client) *cobra.Command {
	cc := &cobra.Command{
		Use:   "create",
		Short: "Create a MksTaskRun in default namespace",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Println("MksTaskRun called")
			fs, _ := cmd.Flags().GetString("taskRef")
			if fs == "" {
				fmt.Errorf("taskRef not defined")
			}
			cmtr := &v1alpha1.MksTaskRun{
				TypeMeta:   metav1.TypeMeta{Kind: "MksTaskRun"},
				ObjectMeta: metav1.ObjectMeta{Name: "mkstaskrun" + fs},
				Spec: v1alpha1.MksTaskRunSpec{
					TaskRef: v1alpha1.MksTaskRef{
						Name: fs,
					},
				},
			}
			obj, err := mksc.Mks.MkscontrollerV1alpha1().MksTaskRuns("default").Create(context.TODO(), cmtr, metav1.CreateOptions{})
			if err != nil {
				return nil
			}
			fmt.Println(obj)
			return nil
		},
	}
	cc.Flags().String("taskRef", "", "Task name to be executed")
	return cc
}
