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
	"testing"

	"github.com/MiniTeks/mks-cli/pkg/test"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestCreate(t *testing.T) {
	fc := &test.FakeMksParam{}
	fc.SetNamespace("default")
	fc.ClearObjects()
	fc.SetTestObjects()
	cs, _ := fc.Client(nil)

	tr := Command(cs)
	_, err := test.ExecuteCommand(tr, "create", "--taskRef=testcreate")
	obj, err := cs.Mks.MkscontrollerV1alpha1().MksTaskRuns("default").Get(context.Background(), "mkstaskruntestcreate", v1.GetOptions{})

	if err != nil {
		t.Fatalf("Cannot execute command: %v", err)
	} else if obj.Spec.TaskRef.Name != "testcreate" {
		t.Fatal("Cant create taskrun")
	}
}
