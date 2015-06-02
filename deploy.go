// Copyright (c) 2015 Pagoda Box Inc
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v.
// 2.0. If a copy of the MPL was not distributed with this file, You can obtain one
// at http://mozilla.org/MPL/2.0/.
//

package client

import (
	"encoding/json"
)

//
type (

	// Deploy represents a Pagoda Box app deploy
	Deploy struct { }

	// DeployCreateOptions represents all available options when creating an app.
	DeployCreateOptions struct { }
)

// routes

// CreateDeploy creates a new app, with provided options
func CreateDeploy(options *DeployCreateOptions) (*Deploy, error) {

	b, err := json.Marshal(options)
	if err != nil {
		return nil, err
	}

	var deploy Deploy
	return &deploy, post(&deploy, "/apps", string(b))
}
