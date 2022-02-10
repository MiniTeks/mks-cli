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

	"github.com/MiniTeks/mks-server/pkg/apis/mkscontroller/v1alpha1"
	"github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
)

var myFlags fl = fl{}

func MksTaskCreate(mksclient *versioned.Clientset) *cobra.Command {
	mksTaskCreateCmd := &cobra.Command{
		Use:   "create",
		Short: "ceates mkstask",
		Long:  "mkstask create is used to create",
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
				Spec:       v1alpha1.MksTaskSpec{Name: myFlags.stepname, Image: myFlags.image, Command: myFlags.command, Args: myFlags.args},
			}
			namespace, _ := cmd.Flags().GetString("namespace")
			crt, err := mksclient.MkscontrollerV1alpha1().MksTasks(myFlags.namespace).Create(context.TODO(), mt, v1.CreateOptions{})
			if err != nil {
				fmt.Printf("Error!!! Coldn't create the resource with name %s in the namespace %s\n", name, namespace)
				fmt.Errorf("Couldn't create mksTsk", err.Error())
				return err
			} else {
				fmt.Println(string(crt.UID))
			}
			fmt.Println(crt)
			fmt.Println(args, "mkstask create called")
			return nil
		},
	}
	mksTaskCreateCmd.Flags().StringVarP(&myFlags.namespace, "namespace", "n", "default", "namespace of the mksTaskResource")
	mksTaskCreateCmd.Flags().StringVarP(&myFlags.stepname, "stepname", "s", "", "provide step name")
	mksTaskCreateCmd.MarkFlagRequired("sn")
	mksTaskCreateCmd.Flags().StringVarP(&myFlags.image, "image", "i", "", "name of the image resource")
	mksTaskCreateCmd.MarkFlagRequired("i")
	mksTaskCreateCmd.Flags().StringVarP(&myFlags.command, "command", "c", "", "commands")
	mksTaskCreateCmd.MarkFlagRequired("command")
	mksTaskCreateCmd.Flags().StringVarP(&myFlags.args, "args", "a", "", "arguments to the commmand")
	mksTaskCreateCmd.MarkFlagRequired("args")
	return mksTaskCreateCmd
}
