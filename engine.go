package client

import (
	"encoding/json"
	"time"
)

//
type (

	// Engine represents a published nanobox engine
	Engine struct {
		EngineConfig
		ActiveReleaseID string    `json:"active_release_id"`
		CreatedAt       time.Time `json:"created_at"`
		HasIcon         bool      `json:"has_icon"`
		ID              string    `json:"id"`
		Official        bool      `json:"official"`
		Private         bool      `json:"private"`
		StarCount       string    `json:"star_count"`
		State           string    `json:"state"`
		UpdatedAt       time.Time `json:"updated_at"`
	}

	// EngineConfig represents all available options when creating an engine
	EngineConfig struct {
		Name string `json:"app_id,omitempty"`
	}
)

// CreateEngine creates a new engine
func CreateEngine(config *EngineConfig) (*Engine, error) {

	//
	b, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	//
	var engine Engine
	return &engine, post(&engine, "/engines", string(b))
}

// GetEngine returns the specified engine
func GetEngine(userSlug, engineSlug string) (*Engine, error) {

	var path string

	switch {
	case userSlug == "":
		path = "/engines/" + engineSlug
	default:
		path = "/engines/" + userSlug + "/" + engineSlug
	}

	var engine Engine
	return &engine, get(&engine, path)
}
