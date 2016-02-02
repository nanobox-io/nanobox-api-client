package client

import (
	"encoding/json"
	"time"
)

//
type (

	// Service represents a published nanobox service
	Service struct {
		ServiceConfig
		CreatedAt time.Time `json:"created_at"`
		HasIcon   bool      `json:"has_icon"`
		ID        string    `json:"id"`
		Official  bool      `json:"official"`
		Private   bool      `json:"private"`
		StarCount string    `json:"star_count"`
		State     string    `json:"state"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	// ServiceConfig represents all available options when creating an service
	ServiceConfig struct {
		Authors     *[]string `json:"authors,omitempty"`
		Behaviors   *[]string `json:"behaviors,omitempty"`
		Category    *string   `json:"category,omitempty"`
		Image       *string   `json:"image,omitempty"`
		License     *string   `json:"license,omitempty"`
		Links       *string   `json:"links,omitempty"`
		Name        *string   `json:"name,omitempty"`
		Readme      *string   `json:"readme,omitempty"`
		Stabilities *[]string `json:"stabilities,omitempty"`
		Summary     *string   `json:"summary,omitempty"`
		Topologies  *[]string `json:"topologies,omitempty"`
		Versions    *[]string `json:"versions,omitempty"`
	}
)

// routes

// CreateService creates a new service
func CreateService(config *ServiceConfig) (*Service, error) {

	//
	b, err := json.Marshal(config)
	if err != nil {
		return nil, err
	}

	//
	var service Service
	return &service, post(&service, "/services", string(b))
}

// GetService returns the specified service
func GetService(userSlug, engineSlug, serviceSlug string) (*Service, error) {

	var path string

	switch {
	case userSlug == "":
		path = "/engines/" + engineSlug + "/services/" + serviceSlug
	default:
		path = "/engines/" + userSlug + "/" + engineSlug + "/services/" + serviceSlug
	}

	var service Service
	return &service, get(&service, path)
}
