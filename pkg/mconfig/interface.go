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

package mconfig

import (
	"github.com/MiniTeks/mks-server/pkg/client/clientset/versioned"
	"k8s.io/client-go/rest"
)

type MksParams struct {
	client    *Client
	namespace string
}

var _ Params = (*MksParams)(nil)

func (p *MksParams) Client(config *rest.Config) (*Client, error) {
	if p.client != nil {
		return p.client, nil
	}

	mks, err := p.mksClient(config)
	if err != nil {
		return nil, err
	}

	p.client = &Client{
		Mks: mks,
	}

	return p.client, nil

}

func (p *MksParams) SetNamespace(ns string) {
	p.namespace = ns
}

func (p *MksParams) Namespace() string {
	return p.namespace
}

func (p *MksParams) mksClient(config *rest.Config) (versioned.Interface, error) {
	cfg, err := versioned.NewForConfig(config)
	if err != nil {
		return nil, err
	}
	return cfg, err
}
