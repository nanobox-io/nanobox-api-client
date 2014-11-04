package client

import (
	"time"
)

//
type (

	// AppService represents a Pagoda Box app service
	AppService struct {
		ActivePlanID string            `json:"active_plan_id"` //
		AppID        string            `json:"app_id"`         //
		Authable     bool              `json:"authable"`       //
		Backupable   bool              `json:"backupable"`     //
		Buildable    bool              `json:"buildable"`      //
		Category     string            `json:"category"`       //
		Codeable     bool              `json:"codeable"`       //
		CreatedAt    time.Time         `json:"created_at"`     //
		Dependence   string            `json:"dependence"`     //
		Execable     bool              `json:"execable"`       //
		Host         string            `json:"host"`           //
		ID           string            `json:"id"`             //
		IPs          map[string]string `json:"ips"`            //
		Libable      bool              `json:"libable"`        //
		Name         string            `json:"name"`           //
		Passwords    map[string]string `json:"passwords"`      //
		PublicTunnel bool              `json:"public_tunnel"`  //
		Pulse        string            `json:"pulse"`          //
		Rebootable   bool              `json:"rebootable"`     //
		ScaffoldID   string            `json:"scaffold_id"`    //
		State        string            `json:"state"`          //
		StaticName   string            `json:"static_name"`    //
		Topology     string            `json:"topology"`       //
		TunnelIP     string            `json:"tunnel_ip"`      //
		TunnelPort   int               `json:"tunnel_port"`    //
		TunnelUser   string            `json:"tunnel_user"`    //
		UID          string            `json:"uid"`            //
		UpdatedAt    time.Time         `json:"updated_at"`     //
		Usernames    map[string]string `json:"usernames"`      //
	}

	// AppSerivceCreateOptions represents all available options when creating an
	// app service
	AppServiceCreateOptions struct {
		AppID      string `json:"app_id,omitempty"`
		ExecCmd    string `json:"exec_cmd,omitempty"` //
		Name       string `json:"name,omitempty"`
		ScaffoldID string `json:"scaffold_id,omitempty"`
		Topology   string `json:"topology,omitempty"`
	}

	// AppSerivceUpdateOptions represents all available options when updating an
	// app service
	AppServiceUpdateOptions struct {
		ExecCmd      string `json:"exec_cmd,omitempty"`
		Name         string `json:"name,omitempty"`
		PublicTunnel bool   `json:"public_tunnel,omitempty"`
	}
)

// routes

// GetAppServices returns an index of all an app's services
func (c *Client) GetAppServices(appSlug string) ([]AppService, error) {
	var appServices []AppService
	return appServices, c.get(&appServices, "/apps/"+appSlug+"/services")
}

// CreateAppService creates a new app service, with provided options
func (c *Client) CreateAppService(appSlug string, options *AppServiceCreateOptions) (*AppService, error) {

	body := toJSON(options)

	var appService AppService
	return &appService, c.post(&appService, "/apps/"+appSlug+"/services", string(body))
}

// GetAppService returns the specified app service
func (c *Client) GetAppService(appSlug, appServiceID string) (*AppService, error) {
	var appService AppService
	return &appService, c.get(&appService, "/apps/"+appSlug+"/services/"+appServiceID)
}

// UpdateAppService updates the specified app service, with provided options
func (c *Client) UpdateAppService(appSlug, appServiceID string, options *AppServiceUpdateOptions) (*AppService, error) {

	body := toJSON(options)

	var appService AppService
	return &appService, c.put(&appService, "/apps/"+appSlug+"/services/"+appServiceID, string(body))
}

// DeleteAppService deletes the specified app service
func (c *Client) DeleteAppService(appSlug, appServiceID string) error {
	return c.delete("/apps/" + appSlug + "/services/" + appServiceID)
}

// additional routes

// RebootAppService will reboot a specified app service
func (c *Client) RebootAppService(appSlug, appServiceID string) (*AppService, error) {
	return nil, c.put(nil, "/apps/"+appSlug+"/services/"+appServiceID+"/reboot", nil)
}

// RepairAppService will reboot a specified app service
func (c *Client) RepairAppService(appSlug, appServiceID string) (*AppService, error) {
	return nil, c.put(nil, "/apps/"+appSlug+"/services/"+appServiceID+"/repair", nil)
}

// RestartAppService will reboot a specified app service
func (c *Client) RestartAppService(appSlug, appServiceID string) (*AppService, error) {
	return nil, c.put(nil, "/apps/"+appSlug+"/services/"+appServiceID+"/restart", nil)
}
