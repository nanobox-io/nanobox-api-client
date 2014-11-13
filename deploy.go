package client

import (
	"errors"
	"time"
)

//
type (

	// AppDeploy represents a Nanobox app deploy
	AppDeploy struct {
		AppID        string    `json:"app_id"`        //
		BuildState   string    `json:"build_state"`   //
		BuildSuccess bool      `json:"build_success"` //
		Commit       string    `json:"commit"`        //
		CreatedAt    int       `json:"created_at"`    //
		Email        string    `json:"email"`         //
		GitBranch    string    `json:"git_branch"`    //
		ID           string    `json:"id"`            //
		Message      string    `json:"message"`       //
		State        string    `json:"state"`         //
		UpdatedAt    time.Time `json:"updated_at"`    //
		Username     string    `json:"username"`      //
	}

	// AppDeployCreateOptions represents all available options when creating an app
	// deploy
	AppDeployCreateOptions struct{}

	// AppDeployUpdateOptions represents all available options when updating an app
	// deploy
	AppDeployUpdateOptions struct{}
)

// routes

// GetAppDeploys returns an index of all an app's deploys
func (c *Client) GetAppDeploys(appSlug string) ([]AppDeploy, error) {
	var appDeploys []AppDeploy
	return appDeploys, c.get(&appDeploys, "/apps/"+appSlug+"/deploys")
}

// CreateAppDeploy creates a new app deploy, with provided options
func (c *Client) CreateAppDeploy(appSlug string, options *AppDeployCreateOptions) (*AppDeploy, error) {
	return nil, errors.New("Creating an app deploy is not allowed.")
}

// GetAppDeploy returns the specified app deploy
func (c *Client) GetAppDeploy(appSlug, appDeployID string) (*AppDeploy, error) {
	var appDeploy AppDeploy
	return &appDeploy, c.get(&appDeploy, "/apps/"+appSlug+"/deploys/"+appDeployID)
}

// UpdateAppDeploy updates the specified app deploy, with provided options
func (c *Client) UpdateAppDeploy(appSlug, appDeployID string, options *AppDeployUpdateOptions) (*AppDeploy, error) {
	return nil, errors.New("Updating an app deploy is not allowed.")
}

// DeleteAppDeploy deletes the specified app deploy
func (c *Client) DeleteAppDeploy(appSlug, appDeployID string) error {
	return errors.New("Deleting an app deploy is not allowed.")
}

// additional routes

// RedployAppDeploy redeploys the specified app deploy
func (c *Client) RedeployAppDeploy(appSlug, appDeployID string) (*AppDeploy, error) {
	return nil, c.put(nil, "/apps/"+appSlug+"/deploys/"+appDeployID+"/redeploy", nil)
}
