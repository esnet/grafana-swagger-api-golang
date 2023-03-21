// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/esnet/grafana-swagger-api-golang/goclient/client/access_control"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/access_control_provisioning"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/admin"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/admin_ldap"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/admin_provisioning"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/admin_users"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/annotations"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/api_keys"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/correlations"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/dashboard_permissions"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/dashboard_versions"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/dashboards"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/datasource_permissions"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/datasources"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/ds"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/folder_permissions"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/folders"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/get_current_org"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/ldap_debug"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/legacy_alerts"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/legacy_alerts_notification_channels"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/library_elements"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/licensing"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/org"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/org_invites"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/org_preferences"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/orgs"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/playlists"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/provisioning"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/query_history"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/recording_rules"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/reports"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/saml"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/search"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/service_accounts"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/signed_in_user"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/snapshots"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/sync_team_groups"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/teams"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/user_preferences"
	"github.com/esnet/grafana-swagger-api-golang/goclient/client/users"
)

// Default grafana HTTP API HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/api"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"http", "https"}

// NewHTTPClient creates a new grafana HTTP API HTTP client.
func NewHTTPClient(formats strfmt.Registry) *GrafanaHTTPAPI {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new grafana HTTP API HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *GrafanaHTTPAPI {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new grafana HTTP API client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *GrafanaHTTPAPI {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(GrafanaHTTPAPI)
	cli.Transport = transport
	cli.AccessControl = access_control.New(transport, formats)
	cli.AccessControlProvisioning = access_control_provisioning.New(transport, formats)
	cli.Admin = admin.New(transport, formats)
	cli.AdminLdap = admin_ldap.New(transport, formats)
	cli.AdminProvisioning = admin_provisioning.New(transport, formats)
	cli.AdminUsers = admin_users.New(transport, formats)
	cli.Annotations = annotations.New(transport, formats)
	cli.APIKeys = api_keys.New(transport, formats)
	cli.Correlations = correlations.New(transport, formats)
	cli.DashboardPermissions = dashboard_permissions.New(transport, formats)
	cli.DashboardVersions = dashboard_versions.New(transport, formats)
	cli.Dashboards = dashboards.New(transport, formats)
	cli.DatasourcePermissions = datasource_permissions.New(transport, formats)
	cli.Datasources = datasources.New(transport, formats)
	cli.Ds = ds.New(transport, formats)
	cli.FolderPermissions = folder_permissions.New(transport, formats)
	cli.Folders = folders.New(transport, formats)
	cli.GetCurrentOrg = get_current_org.New(transport, formats)
	cli.LdapDebug = ldap_debug.New(transport, formats)
	cli.LegacyAlerts = legacy_alerts.New(transport, formats)
	cli.LegacyAlertsNotificationChannels = legacy_alerts_notification_channels.New(transport, formats)
	cli.LibraryElements = library_elements.New(transport, formats)
	cli.Licensing = licensing.New(transport, formats)
	cli.Org = org.New(transport, formats)
	cli.OrgInvites = org_invites.New(transport, formats)
	cli.OrgPreferences = org_preferences.New(transport, formats)
	cli.Orgs = orgs.New(transport, formats)
	cli.Playlists = playlists.New(transport, formats)
	cli.Provisioning = provisioning.New(transport, formats)
	cli.QueryHistory = query_history.New(transport, formats)
	cli.RecordingRules = recording_rules.New(transport, formats)
	cli.Reports = reports.New(transport, formats)
	cli.Saml = saml.New(transport, formats)
	cli.Search = search.New(transport, formats)
	cli.ServiceAccounts = service_accounts.New(transport, formats)
	cli.SignedInUser = signed_in_user.New(transport, formats)
	cli.Snapshots = snapshots.New(transport, formats)
	cli.SyncTeamGroups = sync_team_groups.New(transport, formats)
	cli.Teams = teams.New(transport, formats)
	cli.UserPreferences = user_preferences.New(transport, formats)
	cli.Users = users.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// GrafanaHTTPAPI is a client for grafana HTTP API
type GrafanaHTTPAPI struct {
	AccessControl access_control.ClientService

	AccessControlProvisioning access_control_provisioning.ClientService

	Admin admin.ClientService

	AdminLdap admin_ldap.ClientService

	AdminProvisioning admin_provisioning.ClientService

	AdminUsers admin_users.ClientService

	Annotations annotations.ClientService

	APIKeys api_keys.ClientService

	Correlations correlations.ClientService

	DashboardPermissions dashboard_permissions.ClientService

	DashboardVersions dashboard_versions.ClientService

	Dashboards dashboards.ClientService

	DatasourcePermissions datasource_permissions.ClientService

	Datasources datasources.ClientService

	Ds ds.ClientService

	FolderPermissions folder_permissions.ClientService

	Folders folders.ClientService

	GetCurrentOrg get_current_org.ClientService

	LdapDebug ldap_debug.ClientService

	LegacyAlerts legacy_alerts.ClientService

	LegacyAlertsNotificationChannels legacy_alerts_notification_channels.ClientService

	LibraryElements library_elements.ClientService

	Licensing licensing.ClientService

	Org org.ClientService

	OrgInvites org_invites.ClientService

	OrgPreferences org_preferences.ClientService

	Orgs orgs.ClientService

	Playlists playlists.ClientService

	Provisioning provisioning.ClientService

	QueryHistory query_history.ClientService

	RecordingRules recording_rules.ClientService

	Reports reports.ClientService

	Saml saml.ClientService

	Search search.ClientService

	ServiceAccounts service_accounts.ClientService

	SignedInUser signed_in_user.ClientService

	Snapshots snapshots.ClientService

	SyncTeamGroups sync_team_groups.ClientService

	Teams teams.ClientService

	UserPreferences user_preferences.ClientService

	Users users.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *GrafanaHTTPAPI) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.AccessControl.SetTransport(transport)
	c.AccessControlProvisioning.SetTransport(transport)
	c.Admin.SetTransport(transport)
	c.AdminLdap.SetTransport(transport)
	c.AdminProvisioning.SetTransport(transport)
	c.AdminUsers.SetTransport(transport)
	c.Annotations.SetTransport(transport)
	c.APIKeys.SetTransport(transport)
	c.Correlations.SetTransport(transport)
	c.DashboardPermissions.SetTransport(transport)
	c.DashboardVersions.SetTransport(transport)
	c.Dashboards.SetTransport(transport)
	c.DatasourcePermissions.SetTransport(transport)
	c.Datasources.SetTransport(transport)
	c.Ds.SetTransport(transport)
	c.FolderPermissions.SetTransport(transport)
	c.Folders.SetTransport(transport)
	c.GetCurrentOrg.SetTransport(transport)
	c.LdapDebug.SetTransport(transport)
	c.LegacyAlerts.SetTransport(transport)
	c.LegacyAlertsNotificationChannels.SetTransport(transport)
	c.LibraryElements.SetTransport(transport)
	c.Licensing.SetTransport(transport)
	c.Org.SetTransport(transport)
	c.OrgInvites.SetTransport(transport)
	c.OrgPreferences.SetTransport(transport)
	c.Orgs.SetTransport(transport)
	c.Playlists.SetTransport(transport)
	c.Provisioning.SetTransport(transport)
	c.QueryHistory.SetTransport(transport)
	c.RecordingRules.SetTransport(transport)
	c.Reports.SetTransport(transport)
	c.Saml.SetTransport(transport)
	c.Search.SetTransport(transport)
	c.ServiceAccounts.SetTransport(transport)
	c.SignedInUser.SetTransport(transport)
	c.Snapshots.SetTransport(transport)
	c.SyncTeamGroups.SetTransport(transport)
	c.Teams.SetTransport(transport)
	c.UserPreferences.SetTransport(transport)
	c.Users.SetTransport(transport)
}
