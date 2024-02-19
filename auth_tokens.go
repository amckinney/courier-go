// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/trycourier/courier-go/v3/core"
)

type IssueTokenParams struct {
	Scope     string `json:"scope" url:"scope"`
	ExpiresIn string `json:"expires_in" url:"expires_in"`
}

type IssueTokenResponse struct {
	Token *string `json:"token,omitempty" url:"token,omitempty"`

	_rawJSON json.RawMessage
}

func (i *IssueTokenResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler IssueTokenResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*i = IssueTokenResponse(value)
	i._rawJSON = json.RawMessage(data)
	return nil
}

func (i *IssueTokenResponse) String() string {
	if len(i._rawJSON) > 0 {
		if value, err := core.StringifyJSON(i._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(i); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", i)
}