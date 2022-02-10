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

	"github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	"github.com/spf13/cobra"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/klog"
)

func MksTaskDelete(mksclient *versioned.Clientset) *cobra.Command {
	mksTaskDeleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "deletes mkstask",
		Long:  "mkstask delete is used to delete resource",
		RunE: func(cmd *cobra.Command, args []string) error {
			var name string = ""
			if len(args) == 0 {
				klog.Fatalf("Name argument is required to identify your resource")
			} else {
				name = args[0]
			}
			namespace, _ := cmd.Flags().GetString("namespace")
			er := mksclient.MkscontrollerV1alpha1().MksTasks(namespace).Delete(context.TODO(), name, v1.DeleteOptions{})
			if er != nil {
				fmt.Printf("Error!!! Coldn't delete the resource with name %s from the namespace %s\n", name, namespace)
				fmt.Errorf("Couldn't delete mksTsk", er.Error())
				return er
			} else {
				fmt.Println("Successively deleted")
			}
			return nil
		},
	}
	mksTaskDeleteCmd.Flags().StringP("namespace", "n", "default", "namespace of the mksTaskResource")

	return mksTaskDeleteCmd
}
