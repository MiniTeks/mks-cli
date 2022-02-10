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
)

func displayCrtList(crt *v1alpha1.MksTaskList) {
	for _, item := range crt.Items {
		fmt.Print("\n")
		displayCrt(&item)
	}
}

func MksTaskList(mksclient *versioned.Clientset) *cobra.Command {
	mksTaskListCmd := &cobra.Command{
		Use:   "list",
		Short: "lists mkstasks",
		Long:  "mkstask list is used to list",
		RunE: func(cmd *cobra.Command, args []string) error {
			namespace, _ := cmd.Flags().GetString("namespace")
			crt, err := mksclient.MkscontrollerV1alpha1().MksTasks(namespace).List(context.TODO(), v1.ListOptions{})
			if err != nil {
				fmt.Printf("Error!!! Coldn't get the resource(s) from the namespace %s\n", namespace)
				fmt.Errorf("Couldn't create mksTsk", err.Error())
				return err
			}
			fmt.Printf("Here are resources in the namespace %s\n", namespace)
			displayCrtList(crt)
			return nil
		},
	}
	mksTaskListCmd.Flags().StringP("namespace", "n", "default", "namespace of the mksTaskResource")

	return mksTaskListCmd
}
