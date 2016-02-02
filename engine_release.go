package client

import (
	"encoding/json"
	"time"
)

//
type (

	// EngineRelease represents a published nanobox engine release
	EngineRelease struct {
		EngineReleaseConfig
		Checksum  string    `json:"checksum"`
		CreatedAt time.Time `json:"created_at"`
		ID        string    `json:"id"`
		State     string    `json:"state"`
		UpdatedAt time.Time `json:"updated_at"`
		UUID      string    `json:"uuid"`
	}

	// EngineReleaseConfig represents all available options when creating an engine
	// release
	EngineReleaseConfig struct {
		Authors   []string `json:"authors,omitempty"`
		License   string   `json:"license,omitempty"`
		Name      string   `json:"name,omitempty"`
		Readme    string   `json:"readme,omitempty"`
		Stability string   `json:"stability,omitempty"`
		Summary   string   `json:"summary,omitempty"`
		Version   string   `json:"version,omitempty"`
	}
)

// CreateEngineRelease creates a new engine
func CreateEngineRelease(engineSlug string, config *EngineReleaseConfig) (*EngineRelease, error) {

	//
	b, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	//
	var release EngineRelease
	return &release, post(&release, "/engines/"+engineSlug+"/releases", string(b))
}
