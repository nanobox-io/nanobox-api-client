// Copyright (c) 2015 Pagoda Box Inc
//
// This Source Code Form is subject to the terms of the Mozilla Public License, v.
// 2.0. If a copy of the MPL was not distributed with this file, You can obtain one
// at http://mozilla.org/MPL/2.0/.
//

package client

import (
	"encoding/json"
	"time"
)

//
type (

	// EngineRelease represents a nanobox published project
	EngineRelease struct {
		Authors   []string  `json:"authors"`
		Checksum  string    `json:"checksum"`
		CreatedAt time.Time `json:"created_at"`
		Generic   string    `json:"generic"`
		ID        string    `json:"id"`
		Language  string    `json:"language"`
		License   string    `json:"license"`
		Name      string    `json:"name"`
		Readme    string    `json:"readme"`
		Stability string    `json:"stability"`
		State     string    `json:"state"`
		Summary   string    `json:"summary"`
		UpdatedAt time.Time `json:"updated_at"`
		UUID      string    `json:"uuid"`
		Version   string    `json:"version"`
	}

	// EngineReleaseCreateOptions represents all available options when creating a release.
	EngineReleaseCreateOptions struct {
		Authors   []string `json:"authors"`
		Checksum  string   `json:"checksum"`
		Generic   string   `json:"generic"`
		Language  string   `json:"language"`
		License   string   `json:"license"`
		Name      string   `json:"name"`
		Readme    string   `json:"readme"`
		Stability string   `json:"stability"`
		State     string   `json:"state"`
		Summary   string   `json:"summary"`
		Version   string   `json:"version"`
	}
)

// routes

// CreateEngineRelease creates a new release, with provided options
func CreateEngineRelease(engineSlug string, options *EngineReleaseCreateOptions) (*EngineRelease, error) {

	b, err := json.Marshal(options)
	if err != nil {
		return nil, err
	}

	var release EngineRelease
	return &release, post(&release, "/engines/"+engineSlug+"/releases/", string(b))
}

// GetEngineRelease returns the specified release
// func GetEngineRelease(engineSlug, releaseSlug string) (*EngineRelease, error) {
// 	var release EngineRelease
// 	return &release, get(&release, "/engines/" + engineSlug + "/releases/" + releaseSlug)
// }
