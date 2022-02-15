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

package test

import (
	"github.com/MiniTeks/mks-cli/pkg/mconfig"
	"github.com/MiniTeks/mks-server/pkg/client/clientset/versioned/fake"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"
)

type FakeMksParam struct {
	fakeclient *mconfig.Client
	namespace  string
	objects    []runtime.Object
}

var _ mconfig.Params = (*FakeMksParam)(nil)

func (p *FakeMksParam) Client(config *rest.Config) (*mconfig.Client, error) {
	if p.fakeclient != nil {
		return p.fakeclient, nil
	}
	if p.objects != nil {
		p.fakeclient = &mconfig.Client{Mks: fake.NewSimpleClientset(p.objects...)}
		return p.fakeclient, nil
	}
	p.fakeclient = &mconfig.Client{Mks: fake.NewSimpleClientset()}
	return p.fakeclient, nil
}

func (p *FakeMksParam) Namespace() string {
	return p.namespace
}

func (p *FakeMksParam) SetNamespace(ns string) {
	p.namespace = ns
}

func (p *FakeMksParam) SetTestObjects(objects ...runtime.Object) {
	p.objects = objects
}
