// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.11.0 DO NOT EDIT.
package api

import (
	"time"
)

const (
	BearerAuthScopes = "BearerAuth.Scopes"
)

// Defines values for GroupPatchOperationOp.
const (
	GroupPatchOperationOpAdd     GroupPatchOperationOp = "add"
	GroupPatchOperationOpRemove  GroupPatchOperationOp = "remove"
	GroupPatchOperationOpReplace GroupPatchOperationOp = "replace"
)

// Defines values for GroupPatchOperationPath.
const (
	GroupPatchOperationPathName  GroupPatchOperationPath = "name"
	GroupPatchOperationPathPeers GroupPatchOperationPath = "peers"
)

// Defines values for PatchMinimumOp.
const (
	PatchMinimumOpAdd     PatchMinimumOp = "add"
	PatchMinimumOpRemove  PatchMinimumOp = "remove"
	PatchMinimumOpReplace PatchMinimumOp = "replace"
)

// Defines values for RoutePatchOperationOp.
const (
	RoutePatchOperationOpAdd     RoutePatchOperationOp = "add"
	RoutePatchOperationOpRemove  RoutePatchOperationOp = "remove"
	RoutePatchOperationOpReplace RoutePatchOperationOp = "replace"
)

// Defines values for RoutePatchOperationPath.
const (
	RoutePatchOperationPathDescription RoutePatchOperationPath = "description"
	RoutePatchOperationPathEnabled     RoutePatchOperationPath = "enabled"
	RoutePatchOperationPathMasquerade  RoutePatchOperationPath = "masquerade"
	RoutePatchOperationPathMetric      RoutePatchOperationPath = "metric"
	RoutePatchOperationPathNetwork     RoutePatchOperationPath = "network"
	RoutePatchOperationPathNetworkId   RoutePatchOperationPath = "network_id"
	RoutePatchOperationPathPeer        RoutePatchOperationPath = "peer"
)

// Defines values for RulePatchOperationOp.
const (
	RulePatchOperationOpAdd     RulePatchOperationOp = "add"
	RulePatchOperationOpRemove  RulePatchOperationOp = "remove"
	RulePatchOperationOpReplace RulePatchOperationOp = "replace"
)

// Defines values for RulePatchOperationPath.
const (
	RulePatchOperationPathDescription  RulePatchOperationPath = "description"
	RulePatchOperationPathDestinations RulePatchOperationPath = "destinations"
	RulePatchOperationPathDisabled     RulePatchOperationPath = "disabled"
	RulePatchOperationPathFlow         RulePatchOperationPath = "flow"
	RulePatchOperationPathName         RulePatchOperationPath = "name"
	RulePatchOperationPathSources      RulePatchOperationPath = "sources"
)

// Group defines model for Group.
type Group struct {
	// Group ID
	Id string `json:"id"`

	// Group Name identifier
	Name string `json:"name"`

	// List of peers object
	Peers []PeerMinimum `json:"peers"`

	// Count of peers associated to the group
	PeersCount int `json:"peers_count"`
}

// GroupMinimum defines model for GroupMinimum.
type GroupMinimum struct {
	// Group ID
	Id string `json:"id"`

	// Group Name identifier
	Name string `json:"name"`

	// Count of peers associated to the group
	PeersCount int `json:"peers_count"`
}

// GroupPatchOperation defines model for GroupPatchOperation.
type GroupPatchOperation struct {
	// Patch operation type
	Op GroupPatchOperationOp `json:"op"`

	// Group field to update in form /<field>
	Path GroupPatchOperationPath `json:"path"`

	// Values to be applied
	Value []string `json:"value"`
}

// Patch operation type
type GroupPatchOperationOp string

// Group field to update in form /<field>
type GroupPatchOperationPath string

// PatchMinimum defines model for PatchMinimum.
type PatchMinimum struct {
	// Patch operation type
	Op PatchMinimumOp `json:"op"`

	// Values to be applied
	Value []string `json:"value"`
}

// Patch operation type
type PatchMinimumOp string

// Peer defines model for Peer.
type Peer struct {
	// Provides information of who activated the Peer. User or Setup Key
	ActivatedBy struct {
		Type  string `json:"type"`
		Value string `json:"value"`
	} `json:"activated_by"`

	// Peer to Management connection status
	Connected bool `json:"connected"`

	// Groups that the peer belongs to
	Groups []GroupMinimum `json:"groups"`

	// Peer ID
	Id string `json:"id"`

	// Peer's IP address
	Ip string `json:"ip"`

	// Last time peer connected to Netbird's management service
	LastSeen time.Time `json:"last_seen"`

	// Peer's hostname
	Name string `json:"name"`

	// Peer's operating system and version
	Os string `json:"os"`

	// Indicates whether SSH server is enabled on this peer
	SshEnabled bool `json:"ssh_enabled"`

	// Peer's daemon or cli version
	Version string `json:"version"`
}

// PeerMinimum defines model for PeerMinimum.
type PeerMinimum struct {
	// Peer ID
	Id string `json:"id"`

	// Peer's hostname
	Name string `json:"name"`
}

// Route defines model for Route.
type Route struct {
	// Route description
	Description string `json:"description"`

	// Route status
	Enabled bool `json:"enabled"`

	// Route Id
	Id string `json:"id"`

	// Indicate if peer should masquerade traffic to this route's prefix
	Masquerade bool `json:"masquerade"`

	// Route metric number. Lowest number has higher priority
	Metric int `json:"metric"`

	// Network range in CIDR format
	Network string `json:"network"`

	// Route network identifier, to group HA routes
	NetworkId string `json:"network_id"`

	// Network type indicating if it is IPv4 or IPv6
	NetworkType string `json:"network_type"`

	// Peer Identifier associated with route
	Peer string `json:"peer"`
}

// RoutePatchOperation defines model for RoutePatchOperation.
type RoutePatchOperation struct {
	// Patch operation type
	Op RoutePatchOperationOp `json:"op"`

	// Route field to update in form /<field>
	Path RoutePatchOperationPath `json:"path"`

	// Values to be applied
	Value []string `json:"value"`
}

// Patch operation type
type RoutePatchOperationOp string

// Route field to update in form /<field>
type RoutePatchOperationPath string

// RouteRequest defines model for RouteRequest.
type RouteRequest struct {
	// Route description
	Description string `json:"description"`

	// Route status
	Enabled bool `json:"enabled"`

	// Indicate if peer should masquerade traffic to this route's prefix
	Masquerade bool `json:"masquerade"`

	// Route metric number. Lowest number has higher priority
	Metric int `json:"metric"`

	// Network range in CIDR format
	Network string `json:"network"`

	// Route network identifier, to group HA routes
	NetworkId string `json:"network_id"`

	// Peer Identifier associated with route
	Peer string `json:"peer"`
}

// Rule defines model for Rule.
type Rule struct {
	// Rule friendly description
	Description string `json:"description"`

	// Rule destination groups
	Destinations []GroupMinimum `json:"destinations"`

	// Rules status
	Disabled bool `json:"disabled"`

	// Rule flow, currently, only "bidirect" for bi-directional traffic is accepted
	Flow string `json:"flow"`

	// Rule ID
	Id string `json:"id"`

	// Rule name identifier
	Name string `json:"name"`

	// Rule source groups
	Sources []GroupMinimum `json:"sources"`
}

// RuleMinimum defines model for RuleMinimum.
type RuleMinimum struct {
	// Rule friendly description
	Description string `json:"description"`

	// Rules status
	Disabled bool `json:"disabled"`

	// Rule flow, currently, only "bidirect" for bi-directional traffic is accepted
	Flow string `json:"flow"`

	// Rule name identifier
	Name string `json:"name"`
}

// RulePatchOperation defines model for RulePatchOperation.
type RulePatchOperation struct {
	// Patch operation type
	Op RulePatchOperationOp `json:"op"`

	// Rule field to update in form /<field>
	Path RulePatchOperationPath `json:"path"`

	// Values to be applied
	Value []string `json:"value"`
}

// Patch operation type
type RulePatchOperationOp string

// Rule field to update in form /<field>
type RulePatchOperationPath string

// SetupKey defines model for SetupKey.
type SetupKey struct {
	// Setup Key expiration date
	Expires time.Time `json:"expires"`

	// Setup Key ID
	Id string `json:"id"`

	// Setup Key value
	Key string `json:"key"`

	// Setup key last usage date
	LastUsed time.Time `json:"last_used"`

	// Setup key name identifier
	Name string `json:"name"`

	// Setup key revocation status
	Revoked bool `json:"revoked"`

	// Setup key status, "valid", "overused","expired" or "revoked"
	State string `json:"state"`

	// Setup key type, one-off for single time usage and reusable
	Type string `json:"type"`

	// Usage count of setup key
	UsedTimes int `json:"used_times"`

	// Setup key validity status
	Valid bool `json:"valid"`
}

// SetupKeyRequest defines model for SetupKeyRequest.
type SetupKeyRequest struct {
	// Expiration time in seconds
	ExpiresIn int `json:"expires_in"`

	// Setup Key name
	Name string `json:"name"`

	// Setup key revocation status
	Revoked bool `json:"revoked"`

	// Setup key type, one-off for single time usage and reusable
	Type string `json:"type"`
}

// User defines model for User.
type User struct {
	// User's email address
	Email string `json:"email"`

	// User ID
	Id string `json:"id"`

	// User's name from idp provider
	Name string `json:"name"`

	// User's Netbird account role
	Role string `json:"role"`
}

// PostApiGroupsJSONBody defines parameters for PostApiGroups.
type PostApiGroupsJSONBody struct {
	Name  string    `json:"name"`
	Peers *[]string `json:"peers,omitempty"`
}

// PatchApiGroupsIdJSONBody defines parameters for PatchApiGroupsId.
type PatchApiGroupsIdJSONBody = []GroupPatchOperation

// PutApiGroupsIdJSONBody defines parameters for PutApiGroupsId.
type PutApiGroupsIdJSONBody struct {
	Name  *string   `json:"Name,omitempty"`
	Peers *[]string `json:"Peers,omitempty"`
}

// PutApiPeersIdJSONBody defines parameters for PutApiPeersId.
type PutApiPeersIdJSONBody struct {
	Name       string `json:"name"`
	SshEnabled bool   `json:"ssh_enabled"`
}

// PostApiRoutesJSONBody defines parameters for PostApiRoutes.
type PostApiRoutesJSONBody = RouteRequest

// PatchApiRoutesIdJSONBody defines parameters for PatchApiRoutesId.
type PatchApiRoutesIdJSONBody = []RoutePatchOperation

// PutApiRoutesIdJSONBody defines parameters for PutApiRoutesId.
type PutApiRoutesIdJSONBody = RouteRequest

// PostApiRulesJSONBody defines parameters for PostApiRules.
type PostApiRulesJSONBody struct {
	// Rule friendly description
	Description  string    `json:"description"`
	Destinations *[]string `json:"destinations,omitempty"`

	// Rules status
	Disabled bool `json:"disabled"`

	// Rule flow, currently, only "bidirect" for bi-directional traffic is accepted
	Flow string `json:"flow"`

	// Rule name identifier
	Name    string    `json:"name"`
	Sources *[]string `json:"sources,omitempty"`
}

// PatchApiRulesIdJSONBody defines parameters for PatchApiRulesId.
type PatchApiRulesIdJSONBody = []RulePatchOperation

// PutApiRulesIdJSONBody defines parameters for PutApiRulesId.
type PutApiRulesIdJSONBody struct {
	// Rule friendly description
	Description  string    `json:"description"`
	Destinations *[]string `json:"destinations,omitempty"`

	// Rules status
	Disabled bool `json:"disabled"`

	// Rule flow, currently, only "bidirect" for bi-directional traffic is accepted
	Flow string `json:"flow"`

	// Rule name identifier
	Name    string    `json:"name"`
	Sources *[]string `json:"sources,omitempty"`
}

// PostApiSetupKeysJSONBody defines parameters for PostApiSetupKeys.
type PostApiSetupKeysJSONBody = SetupKeyRequest

// PutApiSetupKeysIdJSONBody defines parameters for PutApiSetupKeysId.
type PutApiSetupKeysIdJSONBody = SetupKeyRequest

// PostApiGroupsJSONRequestBody defines body for PostApiGroups for application/json ContentType.
type PostApiGroupsJSONRequestBody PostApiGroupsJSONBody

// PatchApiGroupsIdJSONRequestBody defines body for PatchApiGroupsId for application/json ContentType.
type PatchApiGroupsIdJSONRequestBody = PatchApiGroupsIdJSONBody

// PutApiGroupsIdJSONRequestBody defines body for PutApiGroupsId for application/json ContentType.
type PutApiGroupsIdJSONRequestBody PutApiGroupsIdJSONBody

// PutApiPeersIdJSONRequestBody defines body for PutApiPeersId for application/json ContentType.
type PutApiPeersIdJSONRequestBody PutApiPeersIdJSONBody

// PostApiRoutesJSONRequestBody defines body for PostApiRoutes for application/json ContentType.
type PostApiRoutesJSONRequestBody = PostApiRoutesJSONBody

// PatchApiRoutesIdJSONRequestBody defines body for PatchApiRoutesId for application/json ContentType.
type PatchApiRoutesIdJSONRequestBody = PatchApiRoutesIdJSONBody

// PutApiRoutesIdJSONRequestBody defines body for PutApiRoutesId for application/json ContentType.
type PutApiRoutesIdJSONRequestBody = PutApiRoutesIdJSONBody

// PostApiRulesJSONRequestBody defines body for PostApiRules for application/json ContentType.
type PostApiRulesJSONRequestBody PostApiRulesJSONBody

// PatchApiRulesIdJSONRequestBody defines body for PatchApiRulesId for application/json ContentType.
type PatchApiRulesIdJSONRequestBody = PatchApiRulesIdJSONBody

// PutApiRulesIdJSONRequestBody defines body for PutApiRulesId for application/json ContentType.
type PutApiRulesIdJSONRequestBody PutApiRulesIdJSONBody

// PostApiSetupKeysJSONRequestBody defines body for PostApiSetupKeys for application/json ContentType.
type PostApiSetupKeysJSONRequestBody = PostApiSetupKeysJSONBody

// PutApiSetupKeysIdJSONRequestBody defines body for PutApiSetupKeysId for application/json ContentType.
type PutApiSetupKeysIdJSONRequestBody = PutApiSetupKeysIdJSONBody
