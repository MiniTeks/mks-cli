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
	"fmt"
	"testing"

	"github.com/MiniTeks/mks-cli/pkg/test"
)

func TestGet(t *testing.T) {
	fc := &test.FakeMksParam{}
	fc.SetNamespace("default")
	fc.ClearObjects()
	fc.SetTestObjects(GetTestData(Trget...)...)
	cs, _ := fc.Client(nil)

	tr := Command(cs)
	out, err := test.ExecuteCommand(tr, "get", "--name=getmtr1")
	fmt.Println(out)
	if err != nil {
		t.Fatalf("Cannot execute command: %v", err)
	} else if out != "name: getmtr1\nnamespace: default\ntaskrunref: gettaskref\n" {
		t.Fatal("Cant find taskrun")
	}
}
