package client

import (
	"time"
)

//
type (

	// App represents a Nanobox application
	App struct {
		ActiveDeployID      string    `json:"active_deploy_id"`      // The ID of the current deploy.
		ActiveTransactionID string    `json:"active_transaction_id"` // The ID of the current running transaction.
		AutoDeploy          bool      `json:"auto_deploy"`           // If the app has auto deploy enabled.
		AutoReconfigure     bool      `json:"auto_reconfigure"`      // If the app has auto reconfigure enabled.
		CreatedAt           time.Time `json:"created_at"`            // The timestamp of when the app was created.
		DeployBranch        string    `json:"deploy_branch"`         // The current deploy branch (always master).
		Flation             string    `json:"flation"`               // The current state of the app (in/deflated).
		Free                bool      `json:"free"`                  // If the app is 'tinker' or 'production'.
		ID                  string    `json:"id"`                    // The ID of the app.
		Name                string    `json:"name"`                  // The name of the app.
		NewRelicID          string    `json:"new_relic_id"`          // The New Relic ID associated with the app.
		PaymentMethodId     string    `json:"payment_method_id"`     // The ID of the apps current payment method.
		PromoCode           string    `json:"promo_code"`            // The promo code used when creating the app.
		PublicKey           string    `json:"public_key"`            // The public SSH key associated with the app.
		SMTPAuth            string    `json:"smtp_auth"`             // The SMTP authentication for sending app mail.
		SMTPDomain          string    `json:"smtp_domain"`           // The SMTP domain for sending app mail
		SMTPHost            string    `json:"smtp_host"`             // The SMTP host for sending app mail.
		SMTPHostname        string    `json:"smtp_hostname"`         // The SMTP hostname for sending app mail.
		SMTPPass            string    `json:"smtp_pass"`             // The SMTP password for sending app mail.
		SMTPPort            string    `json:"smtp_port"`             // The SMTP port for sending app mail.
		SMTPTLS             bool      `json:"smtp_tls"`              // If the SMTP for sending app mail has TLS enabled.
		SMTPUser            string    `json:"smtp_user"`             // The SMTP user for sending app mail.
		State               string    `json:"state"`                 // The state of the app.
		Timezone            string    `json:"timezone"`              // The timezone the app uses.
		UpdatedAt           time.Time `json:"updated_at"`            // The time of the last update to the app.
	}

	// AppCreateOptions represents all available options when creating an app.
	AppCreateOptions struct {
		AutoReconfigure bool   `json:"auto_reconfigure,omitempty"`
		Free            bool   `json:"free,omitempty"`
		Name            string `json:"name,omitempty"`
		PromoCode       string `json:"promo_code,omitempty"`
		QuickstartID    string `json:"quickstart_id,omitempty"` // The ID of the quickstart to use when launching the app
		Timezone        string `json:"timezone,omitempty"`
		// UserID          string `json:"user_id,omitempty"` 			 // The ID of user who will own the app (assumed from auth_token)
	}

	// AppUpdateOptions represents all available options when updating an app.
	AppUpdateOptions struct {
		AutoReconfigure string `json:"auto_reconfigure,omitempty"`
		Name            string `json:"name,omitempty"`
		NewRelicID      string `json:"new_relic_id,omitempty"`
		PaymentMethodID string `json:"payment_method_id,omitempty"`
		PromoCode       string `json:"promo_code,omitempty"`
		SMTPAuth        string `json:"smtp_auth,omitempty"`
		SMTPDomain      string `json:"smtp_domain,omitempty"`
		SMTPFrom        string `json:"smtp_from,omitempty"`
		SMTPHost        string `json:"smtp_host,omitempty"`
		SMTPHostname    string `json:"smtp_hostname,omitempty"`
		SMTPPass        string `json:"smtp_pass,omitempty"`
		SMTPPort        string `json:"smtp_port,omitempty"`
		SMTPTLS         string `json:"smtp_tls,omitempty"`
		SMTPUser        string `json:"smtp_user,omitempty"`
		Timezone        string `json:"timezone,omitempty"`
	}
)

// routes

// GetApps returns an index of all of user's apps
func (c *Client) GetApps() ([]App, error) {
	var apps []App
	return apps, c.get(&apps, "/apps")
}

// CreateApp creates a new app, with provided options
func (c *Client) CreateApp(options *AppCreateOptions) (*App, error) {

	body := toJSON(options)

	var app App
	return &app, c.post(&app, "/apps", string(body))
}

// GetApp returns the specified app
func (c *Client) GetApp(appSlug string) (*App, error) {
	var app App
	return &app, c.get(&app, "/apps/"+appSlug)
}

// UpdateApp updates the specified app, with provided options
func (c *Client) UpdateApp(appSlug string, options *AppUpdateOptions) (*App, error) {

	body := toJSON(options)

	var app App
	return &app, c.put(&app, "/apps/"+appSlug, string(body))
}

// DeleteApp deletes the specified app
func (c *Client) DeleteApp(appSlug string) error {
	return c.delete("/apps/" + appSlug)
}

// additional routes

// RebuildAppLibs rebuilds an apps libraries
func (c *Client) RebuildAppLibs(appSlug string) (*App, error) {
	return nil, c.put(nil, "/apps/"+appSlug+"/rebuild", nil)
}

// CleanAppLibs cleans an apps libraries
func (c *Client) CleanAppLibs(appSlug string) (*App, error) {
	return nil, c.put(nil, "/apps/"+appSlug+"/clean_libs", nil)
}

// CleanAndRebuildAppLibs cleans and rebuilds and apps libraries
func (c *Client) CleanAndRebuildAppLibs(appSlug string) (*App, error) {
	return nil, c.put(nil, "/apps/"+appSlug+"/clean_libs_and_rebuild", nil)
}

// RollbackAppDeploy rolls an app back one (1) previous deploy
func (c *Client) RollbackAppDeploy(appSlug string) (*AppDeploy, error) {
	return nil, c.put(nil, "/apps/"+appSlug+"/rollback", nil)
}
