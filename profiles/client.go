// This file was auto-generated by Fern from our API Definition.

package profiles

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"
	v3 "github.com/trycourier/courier-go/v3"
	core "github.com/trycourier/courier-go/v3/core"
	io "io"
	http "net/http"
	url "net/url"
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

// Returns the specified user profile.
//
// A unique identifier representing the user associated with the requested profile.
func (c *Client) Get(ctx context.Context, userId string) (*v3.ProfileGetResponse, error) {
	baseURL := "https://api.courier.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"profiles/%v", userId)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(v3.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *v3.ProfileGetResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			Headers:      c.header,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Merge the supplied values with an existing profile or create a new profile if one doesn't already exist.
//
// A unique identifier representing the user associated with the requested profile.
func (c *Client) Create(ctx context.Context, userId string, request *v3.MergeProfileRequest) (*v3.MergeProfileResponse, error) {
	baseURL := "https://api.courier.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"profiles/%v", userId)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(v3.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *v3.MergeProfileResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodPost,
			Headers:      c.header,
			Request:      request,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// When using `PUT`, be sure to include all the key-value pairs required by the recipient's profile.
// Any key-value pairs that exist in the profile but fail to be included in the `PUT` request will be
// removed from the profile. Remember, a `PUT` update is a full replacement of the data. For partial updates,
// use the [Patch](https://www.courier.com/docs/reference/profiles/patch/) request.
//
// A unique identifier representing the user associated with the requested profile.
func (c *Client) Replace(ctx context.Context, userId string, request *v3.ReplaceProfileRequest) (*v3.ReplaceProfileResponse, error) {
	baseURL := "https://api.courier.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"profiles/%v", userId)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(v3.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *v3.ReplaceProfileResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodPut,
			Headers:      c.header,
			Request:      request,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Deletes the specified user profile.
//
// A unique identifier representing the user associated with the requested profile.
func (c *Client) Delete(ctx context.Context, userId string) error {
	baseURL := "https://api.courier.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"profiles/%v", userId)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(v3.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodDelete,
			Headers:      c.header,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return err
	}
	return nil
}

// Returns the subscribed lists for a specified user.
//
// A unique identifier representing the user associated with the requested profile.
func (c *Client) GetListSubscriptions(ctx context.Context, userId string, request *v3.GetListSubscriptionsRequest) (*v3.GetListSubscriptionsResponse, error) {
	baseURL := "https://api.courier.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"profiles/%v/lists", userId)

	queryParams := make(url.Values)
	if request.Cursor != nil {
		queryParams.Add("cursor", fmt.Sprintf("%v", *request.Cursor))
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(v3.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *v3.GetListSubscriptionsResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodGet,
			Headers:      c.header,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Subscribes the given user to one or more lists. If the list does not exist, it will be created.
//
// A unique identifier representing the user associated with the requested profile.
func (c *Client) SubscribeToLists(ctx context.Context, userId string, request *v3.SubscribeToListsRequest) (*v3.SubscribeToListsResponse, error) {
	baseURL := "https://api.courier.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"profiles/%v/lists", userId)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(v3.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *v3.SubscribeToListsResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodPost,
			Headers:      c.header,
			Request:      request,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// Removes all list subscriptions for given user.
//
// A unique identifier representing the user associated with the requested profile.
func (c *Client) DeleteListSubscription(ctx context.Context, userId string) (*v3.DeleteListSubscriptionResponse, error) {
	baseURL := "https://api.courier.com"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"profiles/%v/lists", userId)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 400:
			value := new(v3.BadRequestError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return apiError
			}
			return value
		}
		return apiError
	}

	var response *v3.DeleteListSubscriptionResponse
	if err := c.caller.Call(
		ctx,
		&core.CallParams{
			URL:          endpointURL,
			Method:       http.MethodDelete,
			Headers:      c.header,
			Response:     &response,
			ErrorDecoder: errorDecoder,
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
