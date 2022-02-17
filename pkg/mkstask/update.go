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

package mkstask

import (
	"context"
	"fmt"

	"github.com/MiniTeks/mks-cli/pkg/mconfig"
	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
)

var myFlags2 fl = fl{}

func MksTaskUpdate(mksc *mconfig.Client) *cobra.Command {
	mksTaskUpdateCmd := &cobra.Command{
		Use:   "update",
		Short: "updates mkstask resources ",
		Long:  "mkstask update is used to update",
		RunE: func(cmd *cobra.Command, args []string) error {
			var name string = ""
			if len(args) == 0 {
				klog.Fatalf("Name argument is required for your resource")
			} else {
				name = args[0]
			}
			mt := &v1alpha1.MksTask{
				TypeMeta:   v1.TypeMeta{Kind: "MksTask"},
				ObjectMeta: v1.ObjectMeta{Name: name},
				Spec:       v1alpha1.MksTaskSpec{Name: myFlags2.stepname, Image: myFlags2.image, Command: myFlags2.command, Args: myFlags2.args},
			}
			crt, err := mksc.Mks.MkscontrollerV1alpha1().MksTasks(myFlags2.namespace).Update(context.TODO(), mt, v1.UpdateOptions{})
			if err != nil {
				fmt.Errorf("Couldn't update mksTsk: %v", err)
				return err
			}
			fmt.Println(crt)
			fmt.Println(args, "mkstask update called")
			return nil
		},
	}
	mksTaskUpdateCmd.Flags().StringVarP(&myFlags2.namespace, "namespace", "n", "default", "namespace of the mksTaskResource")
	mksTaskUpdateCmd.Flags().StringVarP(&myFlags2.stepname, "stepname", "s", "", "provide step name")
	mksTaskUpdateCmd.MarkFlagRequired("sn")
	mksTaskUpdateCmd.Flags().StringVarP(&myFlags2.image, "image", "i", "", "name of the image resource")
	mksTaskUpdateCmd.MarkFlagRequired("i")
	mksTaskUpdateCmd.Flags().StringVarP(&myFlags2.command, "command", "c", "", "commands")
	mksTaskUpdateCmd.MarkFlagRequired("command")
	mksTaskUpdateCmd.Flags().StringVarP(&myFlags2.args, "args", "a", "", "arguments to the commmand")
	return mksTaskUpdateCmd
}
