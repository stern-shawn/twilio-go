/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Autopilot
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateModelBuild'
type CreateModelBuildParams struct {
	// The URL we should call using a POST method to send status information to your application.
	StatusCallback *string `json:"StatusCallback,omitempty"`
	// An application-defined string that uniquely identifies the new resource. This value must be a unique string of no more than 64 characters. It can be used as an alternative to the `sid` in the URL path to address the resource.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *CreateModelBuildParams) SetStatusCallback(StatusCallback string) *CreateModelBuildParams {
	params.StatusCallback = &StatusCallback
	return params
}
func (params *CreateModelBuildParams) SetUniqueName(UniqueName string) *CreateModelBuildParams {
	params.UniqueName = &UniqueName
	return params
}

//
func (c *ApiService) CreateModelBuild(AssistantSid string, params *CreateModelBuildParams) (*AutopilotV1ModelBuild, error) {
	path := "/v1/Assistants/{AssistantSid}/ModelBuilds"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.StatusCallback != nil {
		data.Set("StatusCallback", *params.StatusCallback)
	}
	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1ModelBuild{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

//
func (c *ApiService) DeleteModelBuild(AssistantSid string, Sid string) error {
	path := "/v1/Assistants/{AssistantSid}/ModelBuilds/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

//
func (c *ApiService) FetchModelBuild(AssistantSid string, Sid string) (*AutopilotV1ModelBuild, error) {
	path := "/v1/Assistants/{AssistantSid}/ModelBuilds/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1ModelBuild{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListModelBuild'
type ListModelBuildParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListModelBuildParams) SetPageSize(PageSize int) *ListModelBuildParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListModelBuildParams) SetLimit(Limit int) *ListModelBuildParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ModelBuild records from the API. Request is executed immediately.
func (c *ApiService) PageModelBuild(AssistantSid string, params *ListModelBuildParams, pageToken, pageNumber string) (*ListModelBuildResponse, error) {
	path := "/v1/Assistants/{AssistantSid}/ModelBuilds"

	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageNumber != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListModelBuildResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ModelBuild records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListModelBuild(AssistantSid string, params *ListModelBuildParams) ([]AutopilotV1ModelBuild, error) {
	response, errors := c.StreamModelBuild(AssistantSid, params)

	records := make([]AutopilotV1ModelBuild, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams ModelBuild records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamModelBuild(AssistantSid string, params *ListModelBuildParams) (chan AutopilotV1ModelBuild, chan error) {
	if params == nil {
		params = &ListModelBuildParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan AutopilotV1ModelBuild, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageModelBuild(AssistantSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamModelBuild(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamModelBuild(response *ListModelBuildResponse, params *ListModelBuildParams, recordChannel chan AutopilotV1ModelBuild, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.ModelBuilds
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListModelBuildResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListModelBuildResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListModelBuildResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListModelBuildResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateModelBuild'
type UpdateModelBuildParams struct {
	// An application-defined string that uniquely identifies the resource. This value must be a unique string of no more than 64 characters. It can be used as an alternative to the `sid` in the URL path to address the resource.
	UniqueName *string `json:"UniqueName,omitempty"`
}

func (params *UpdateModelBuildParams) SetUniqueName(UniqueName string) *UpdateModelBuildParams {
	params.UniqueName = &UniqueName
	return params
}

//
func (c *ApiService) UpdateModelBuild(AssistantSid string, Sid string, params *UpdateModelBuildParams) (*AutopilotV1ModelBuild, error) {
	path := "/v1/Assistants/{AssistantSid}/ModelBuilds/{Sid}"
	path = strings.Replace(path, "{"+"AssistantSid"+"}", AssistantSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.UniqueName != nil {
		data.Set("UniqueName", *params.UniqueName)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &AutopilotV1ModelBuild{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}
