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

func displayCrt(crt *v1alpha1.MksTask) {
	fmt.Println("name : ", crt.Name)
	fmt.Println("namespace : ", crt.Namespace)
	fmt.Println("spec : ", crt.Spec)
}

func MksTaskGet(mksclient *versioned.Clientset) *cobra.Command {
	mksTaskGetCmd := &cobra.Command{
		Use:   "get",
		Short: "gets mkstasks",
		Long:  "mkstask get is used to get",
		RunE: func(cmd *cobra.Command, args []string) error {
			var name string = ""
			if len(args) == 0 {
				klog.Fatalf("A Name argument is required to get your resource")
			} else {
				name = args[0]
			}
			namespace, _ := cmd.Flags().GetString("namespace")
			crt, err := mksclient.MkscontrollerV1alpha1().MksTasks(namespace).Get(context.Background(), name, v1.GetOptions{})
			if err != nil {
				fmt.Printf("Error!!! Coldn't get any resource with name %s inside %s\n", name, namespace)
				klog.Fatal(err.Error())
				return err
			}
			fmt.Println("Here is your requested resource")
			displayCrt(crt)
			return nil

		},
	}
	mksTaskGetCmd.Flags().StringP("namespace", "n", "default", "namespace of the mksTaskResource")

	return mksTaskGetCmd
}
