// This file was auto-generated by Fern from our API Definition.

package automations

import (
	context "context"
	fmt "fmt"
	v3 "github.com/trycourier/courier-go/v3"
	core "github.com/trycourier/courier-go/v3/core"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *core.Caller
	header  http.Header
}

func NewClient(opts ...core.ClientOption) *Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &Client{
		baseURL: options.BaseURL,
		caller:  core.NewCaller(options.HTTPClient),
		header:  options.ToHeader(),
	}
}

// Invoke an automation run from an automation template.
//
// A unique identifier representing the automation template to be invoked. This could be the Automation Template ID or the Automation Template Alias.
func (c *Client) InvokeAutomationTemplate(ctx context.Context, templateId string, request *v3.AutomationInvokeParams) (*v3.AutomationInvokeResponse, error) {
	baseURL := "https://api.courier.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"automations/%v/invoke", templateId)

	var response *v3.AutomationInvokeResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodPost,
			Headers:  c.header,
			Request:  request,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Invoke an ad hoc automation run. This endpoint accepts a JSON payload with a series of automation steps. For information about what steps are available, checkout the ad hoc automation guide [here](https://www.courier.com/docs/automations/steps/).
func (c *Client) InvokeAdHocAutomation(ctx context.Context, request *v3.AutomationAdHocInvokeParams) (*v3.AutomationInvokeResponse, error) {
	baseURL := "https://api.courier.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := baseURL + "/" + "automations/invoke"

	var response *v3.AutomationInvokeResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:      endpointURL,
			Method:   http.MethodPost,
			Headers:  c.header,
			Request:  request,
			Response: &response,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
