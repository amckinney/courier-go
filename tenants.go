// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/trycourier/courier-go/v3/core"
)

type TenantCreateOrReplaceParams struct {
	// Name of the tenant.
	Name string `json:"name"`
	// Tenant's parent id (if any).
	ParentTenantId *string `json:"parent_tenant_id,omitempty"`
	// Defines the preferences used for the tenant when the user hasn't specified their own.
	DefaultPreferences *DefaultPreferences `json:"default_preferences,omitempty"`
	// Arbitrary properties accessible to a template.
	Properties []TemplateProperty `json:"properties,omitempty"`
	// A user profile object merged with user profile on send.
	UserProfile map[string]interface{} `json:"user_profile,omitempty"`
	// Brand to be used for the account when one is not specified by the send call.
	BrandId *string `json:"brand_id,omitempty"`
}

type ListUsersForTenantParams struct {
	// The number of accounts to return
	// (defaults to 20, maximum value of 100)
	Limit *int `json:"-"`
	// Continue the pagination with the next cursor
	Cursor *string `json:"-"`
}

type ListTenantParams struct {
	// The number of accousnts to return
	// (defaults to 20, maximum value of 100)
	Limit *int `json:"-"`
	// Continue the pagination with the next cursor
	Cursor *string `json:"-"`
}

type DefaultPreferences struct {
	Items []*SubscriptionTopic `json:"items,omitempty"`

	_rawJSON json.RawMessage
}

func (d *DefaultPreferences) UnmarshalJSON(data []byte) error {
	type unmarshaler DefaultPreferences
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DefaultPreferences(value)
	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DefaultPreferences) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

type ListUsersForTenantResponse struct {
	Items *UserTenantAssociation `json:"items,omitempty"`
	// Set to true when there are more pages that can be retrieved.
	HasMore bool `json:"has_more"`
	// A url that may be used to generate these results.
	Url string `json:"url"`
	// A url that may be used to generate fetch the next set of results.
	// Defined only when `has_more` is set to true
	NextUrl *string `json:"next_url,omitempty"`
	// A pointer to the next page of results. Defined
	// only when `has_more` is set to true
	Cursor *string `json:"cursor,omitempty"`
	// Always set to `list`. Represents the type of this object.
	type_ string

	_rawJSON json.RawMessage
}

func (l *ListUsersForTenantResponse) Type() string {
	return l.type_
}

func (l *ListUsersForTenantResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListUsersForTenantResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListUsersForTenantResponse(value)
	l.type_ = "list"
	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListUsersForTenantResponse) MarshalJSON() ([]byte, error) {
	type embed ListUsersForTenantResponse
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*l),
		Type:  "list",
	}
	return json.Marshal(marshaler)
}

func (l *ListUsersForTenantResponse) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

type TemplateProperty = interface{}

type Tenant struct {
	// Id of the tenant.
	Id string `json:"id"`
	// Name of the tenant.
	Name string `json:"name"`
	// Tenant's parent id (if any).
	ParentTenantId *string `json:"parent_tenant_id,omitempty"`
	// Defines the preferences used for the account when the user hasn't specified their own.
	DefaultPreferences *DefaultPreferences `json:"default_preferences,omitempty"`
	// Arbitrary properties accessible to a template.
	Properties *TemplateProperty `json:"properties,omitempty"`
	// A user profile object merged with user profile on send.
	UserProfile map[string]interface{} `json:"user_profile,omitempty"`
	// Brand to be used for the account when one is not specified by the send call.
	BrandId *string `json:"brand_id,omitempty"`

	_rawJSON json.RawMessage
}

func (t *Tenant) UnmarshalJSON(data []byte) error {
	type unmarshaler Tenant
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = Tenant(value)
	t._rawJSON = json.RawMessage(data)
	return nil
}

func (t *Tenant) String() string {
	if len(t._rawJSON) > 0 {
		if value, err := core.StringifyJSON(t._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}

type TenantListResponse struct {
	// A pointer to the next page of results. Defined only when has_more is set to true.
	Cursor *string `json:"cursor,omitempty"`
	// Set to true when there are more pages that can be retrieved.
	HasMore bool `json:"has_more"`
	// An array of Tenants
	Items []*Tenant `json:"items,omitempty"`
	// A url that may be used to generate fetch the next set of results.
	// Defined only when has_more is set to true
	NextUrl *string `json:"next_url,omitempty"`
	// A url that may be used to generate these results.
	Url string `json:"url"`
	// Always set to "list". Represents the type of this object.
	type_ string

	_rawJSON json.RawMessage
}

func (t *TenantListResponse) Type() string {
	return t.type_
}

func (t *TenantListResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler TenantListResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*t = TenantListResponse(value)
	t.type_ = "list"
	t._rawJSON = json.RawMessage(data)
	return nil
}

func (t *TenantListResponse) MarshalJSON() ([]byte, error) {
	type embed TenantListResponse
	var marshaler = struct {
		embed
		Type string `json:"type"`
	}{
		embed: embed(*t),
		Type:  "list",
	}
	return json.Marshal(marshaler)
}

func (t *TenantListResponse) String() string {
	if len(t._rawJSON) > 0 {
		if value, err := core.StringifyJSON(t._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(t); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", t)
}
