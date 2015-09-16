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

	// Engine represents a nanobox published project
	Engine struct {
		ActiveReleaseID string    `json:"active_release_id"`
		CreatedAt       time.Time `json:"created_at"`
		CreatorID       string    `json:"creator_id"`
		Downloads       int       `json:"downloads"`
		Generic         bool      `json:"generic"`
		ID              string    `json:"id"`
		LanguageName    string    `json:"language_name"`
		Name            string    `json:"name"`
		Official        bool      `json:"official"`
		State           string    `json:"state"`
		UpdatedAt       time.Time `json:"updated_at"`
	}
)

// routes

// CreateEngine creates a new engine
func CreateEngine(engine *Engine) (*Engine, error) {

	//
	b, err := json.Marshal(engine)
	if err != nil {
		return nil, err
	}

	//
	return engine, post(engine, "/engines", string(b))
}

// GetEngine returns the specified engine
func GetEngine(userSlug, engineSlug string) (*Engine, error) {

	var path string

	//
	path = "/engines/" + userSlug + "/" + engineSlug
	if userSlug == "" {
		path = "/engines/" + engineSlug
	}

	var engine Engine
	return &engine, get(&engine, path)
}
