/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ServerPolicyRuleObservation struct {

	// Lifetime of access token. Can be set to a value between 5 and 1440 minutes. Default is `60`.
	AccessTokenLifetimeMinutes *float64 `json:"accessTokenLifetimeMinutes,omitempty" tf:"access_token_lifetime_minutes,omitempty"`

	// Auth server ID
	AuthServerID *string `json:"authServerId,omitempty" tf:"auth_server_id,omitempty"`

	// Accepted grant type values: authorization_code, implicit, password, client_credentials
	GrantTypeWhitelist []*string `json:"grantTypeWhitelist,omitempty" tf:"grant_type_whitelist,omitempty"`

	// Specifies a set of Groups whose Users are to be excluded.
	GroupBlacklist []*string `json:"groupBlacklist,omitempty" tf:"group_blacklist,omitempty"`

	// Specifies a set of Groups whose Users are to be included. Can be set to Group ID or to the following: `EVERYONE`.
	GroupWhitelist []*string `json:"groupWhitelist,omitempty" tf:"group_whitelist,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The ID of the inline token to trigger.
	InlineHookID *string `json:"inlineHookId,omitempty" tf:"inline_hook_id,omitempty"`

	// Auth server policy rule name
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Auth server policy ID
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Priority of the auth server policy rule
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Lifetime of refresh token.
	RefreshTokenLifetimeMinutes *float64 `json:"refreshTokenLifetimeMinutes,omitempty" tf:"refresh_token_lifetime_minutes,omitempty"`

	// Window in which a refresh token can be used. It can be a value between 5 and 2628000 (5 years) minutes. Default is `10080` (7 days).`refresh_token_window_minutes` must be between `access_token_lifetime_minutes` and `refresh_token_lifetime_minutes`.
	RefreshTokenWindowMinutes *float64 `json:"refreshTokenWindowMinutes,omitempty" tf:"refresh_token_window_minutes,omitempty"`

	// Scopes allowed for this policy rule. They can be whitelisted by name or all can be whitelisted with ` * `
	ScopeWhitelist []*string `json:"scopeWhitelist,omitempty" tf:"scope_whitelist,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// The rule is the system (default) rule for its associated policy
	System *bool `json:"system,omitempty" tf:"system,omitempty"`

	// Auth server policy rule type, unlikely this will be anything other then the default
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies a set of Users to be excluded.
	UserBlacklist []*string `json:"userBlacklist,omitempty" tf:"user_blacklist,omitempty"`

	// Specifies a set of Users to be included.
	UserWhitelist []*string `json:"userWhitelist,omitempty" tf:"user_whitelist,omitempty"`
}

type ServerPolicyRuleParameters struct {

	// Lifetime of access token. Can be set to a value between 5 and 1440 minutes. Default is `60`.
	// +kubebuilder:validation:Optional
	AccessTokenLifetimeMinutes *float64 `json:"accessTokenLifetimeMinutes,omitempty" tf:"access_token_lifetime_minutes,omitempty"`

	// Auth server ID
	// +crossplane:generate:reference:type=github.com/healthcarecom/provider-okta/apis/auth/v1alpha1.Server
	// +crossplane:generate:reference:extractor=github.com/healthcarecom/provider-okta/apis/auth/v1alpha1.AuthServerID()
	// +kubebuilder:validation:Optional
	AuthServerID *string `json:"authServerId,omitempty" tf:"auth_server_id,omitempty"`

	// Reference to a Server in auth to populate authServerId.
	// +kubebuilder:validation:Optional
	AuthServerIDRef *v1.Reference `json:"authServerIdRef,omitempty" tf:"-"`

	// Selector for a Server in auth to populate authServerId.
	// +kubebuilder:validation:Optional
	AuthServerIDSelector *v1.Selector `json:"authServerIdSelector,omitempty" tf:"-"`

	// Accepted grant type values: authorization_code, implicit, password, client_credentials
	// +kubebuilder:validation:Optional
	GrantTypeWhitelist []*string `json:"grantTypeWhitelist,omitempty" tf:"grant_type_whitelist,omitempty"`

	// Specifies a set of Groups whose Users are to be excluded.
	// +kubebuilder:validation:Optional
	GroupBlacklist []*string `json:"groupBlacklist,omitempty" tf:"group_blacklist,omitempty"`

	// Specifies a set of Groups whose Users are to be included. Can be set to Group ID or to the following: `EVERYONE`.
	// +kubebuilder:validation:Optional
	GroupWhitelist []*string `json:"groupWhitelist,omitempty" tf:"group_whitelist,omitempty"`

	// The ID of the inline token to trigger.
	// +kubebuilder:validation:Optional
	InlineHookID *string `json:"inlineHookId,omitempty" tf:"inline_hook_id,omitempty"`

	// Auth server policy rule name
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`

	// Auth server policy ID
	// +crossplane:generate:reference:type=github.com/healthcarecom/provider-okta/apis/auth/v1alpha1.ServerPolicy
	// +crossplane:generate:reference:extractor=github.com/healthcarecom/provider-okta/apis/auth/v1alpha1.AuthServerPolicyId()
	// +kubebuilder:validation:Optional
	PolicyID *string `json:"policyId,omitempty" tf:"policy_id,omitempty"`

	// Reference to a ServerPolicy in auth to populate policyId.
	// +kubebuilder:validation:Optional
	PolicyIDRef *v1.Reference `json:"policyIdRef,omitempty" tf:"-"`

	// Selector for a ServerPolicy in auth to populate policyId.
	// +kubebuilder:validation:Optional
	PolicyIDSelector *v1.Selector `json:"policyIdSelector,omitempty" tf:"-"`

	// Priority of the auth server policy rule
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// Lifetime of refresh token.
	// +kubebuilder:validation:Optional
	RefreshTokenLifetimeMinutes *float64 `json:"refreshTokenLifetimeMinutes,omitempty" tf:"refresh_token_lifetime_minutes,omitempty"`

	// Window in which a refresh token can be used. It can be a value between 5 and 2628000 (5 years) minutes. Default is `10080` (7 days).`refresh_token_window_minutes` must be between `access_token_lifetime_minutes` and `refresh_token_lifetime_minutes`.
	// +kubebuilder:validation:Optional
	RefreshTokenWindowMinutes *float64 `json:"refreshTokenWindowMinutes,omitempty" tf:"refresh_token_window_minutes,omitempty"`

	// Scopes allowed for this policy rule. They can be whitelisted by name or all can be whitelisted with ` * `
	// +kubebuilder:validation:Optional
	ScopeWhitelist []*string `json:"scopeWhitelist,omitempty" tf:"scope_whitelist,omitempty"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Auth server policy rule type, unlikely this will be anything other then the default
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// Specifies a set of Users to be excluded.
	// +kubebuilder:validation:Optional
	UserBlacklist []*string `json:"userBlacklist,omitempty" tf:"user_blacklist,omitempty"`

	// Specifies a set of Users to be included.
	// +kubebuilder:validation:Optional
	UserWhitelist []*string `json:"userWhitelist,omitempty" tf:"user_whitelist,omitempty"`
}

// ServerPolicyRuleSpec defines the desired state of ServerPolicyRule
type ServerPolicyRuleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ServerPolicyRuleParameters `json:"forProvider"`
}

// ServerPolicyRuleStatus defines the observed state of ServerPolicyRule.
type ServerPolicyRuleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ServerPolicyRuleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// ServerPolicyRule is the Schema for the ServerPolicyRules API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,okta}
type ServerPolicyRule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.grantTypeWhitelist)",message="grantTypeWhitelist is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.name)",message="name is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.priority)",message="priority is a required parameter"
	Spec   ServerPolicyRuleSpec   `json:"spec"`
	Status ServerPolicyRuleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ServerPolicyRuleList contains a list of ServerPolicyRules
type ServerPolicyRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ServerPolicyRule `json:"items"`
}

// Repository type metadata.
var (
	ServerPolicyRule_Kind             = "ServerPolicyRule"
	ServerPolicyRule_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: ServerPolicyRule_Kind}.String()
	ServerPolicyRule_KindAPIVersion   = ServerPolicyRule_Kind + "." + CRDGroupVersion.String()
	ServerPolicyRule_GroupVersionKind = CRDGroupVersion.WithKind(ServerPolicyRule_Kind)
)

func init() {
	SchemeBuilder.Register(&ServerPolicyRule{}, &ServerPolicyRuleList{})
}