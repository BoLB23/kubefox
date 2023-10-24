/*
Copyright © 2023 XigXog

This Source Code Form is subject to the terms of the Mozilla Public License,
v2.0. If a copy of the MPL was not distributed with this file, You can obtain
one at https://mozilla.org/MPL/2.0/.
*/

package controller

import (
	"fmt"

	"github.com/xigxog/kubefox/libs/core/kubefox"
)

var (
	ErrNotFound = fmt.Errorf("not found")
)

const (
	NATSImage  = "ghcr.io/xigxog/nats:2.9.23"
	VaultImage = "ghcr.io/xigxog/vault:1.14.4-v0.2.1-alpha"
)

var (
	BrokerImage    = fmt.Sprintf("ghcr.io/xigxog/kubefox/broker:%s", kubefox.Version())
	BootstrapImage = fmt.Sprintf("ghcr.io/xigxog/kubefox/bootstrap:%s", kubefox.Version())
)
