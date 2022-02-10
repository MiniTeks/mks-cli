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
	"github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	"github.com/spf13/cobra"
)

// mkstaskrunCmd represents the mkstaskrun command
var mkstaskrunCmd = &cobra.Command{
	Use:   "mkstaskrun",
	Short: "Add create list mkstaskrun",
}

func Command(mksclient *versioned.Clientset) *cobra.Command {
	mkstaskrunCmd.AddCommand(
		createMksTaskRun(mksclient),
		deleteMksTaskRun(mksclient),
		listMksTaskRun(mksclient),
		getMksTaskRun(mksclient))
	return mkstaskrunCmd
}
