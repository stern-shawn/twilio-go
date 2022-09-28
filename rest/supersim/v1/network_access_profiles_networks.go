/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Supersim
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

// Optional parameters for the method 'CreateNetworkAccessProfileNetwork'
type CreateNetworkAccessProfileNetworkParams struct {
	// The SID of the Network resource to be added to the Network Access Profile resource.
	Network *string `json:"Network,omitempty"`
}

func (params *CreateNetworkAccessProfileNetworkParams) SetNetwork(Network string) *CreateNetworkAccessProfileNetworkParams {
	params.Network = &Network
	return params
}

// Add a Network resource to the Network Access Profile resource.
func (c *ApiService) CreateNetworkAccessProfileNetwork(NetworkAccessProfileSid string, params *CreateNetworkAccessProfileNetworkParams) (*SupersimV1NetworkAccessProfileNetwork, error) {
	path := "/v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks"
	path = strings.Replace(path, "{"+"NetworkAccessProfileSid"+"}", NetworkAccessProfileSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Network != nil {
		data.Set("Network", *params.Network)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1NetworkAccessProfileNetwork{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove a Network resource from the Network Access Profile resource&#39;s.
func (c *ApiService) DeleteNetworkAccessProfileNetwork(NetworkAccessProfileSid string, Sid string) error {
	path := "/v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks/{Sid}"
	path = strings.Replace(path, "{"+"NetworkAccessProfileSid"+"}", NetworkAccessProfileSid, -1)
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

// Fetch a Network Access Profile resource&#39;s Network resource.
func (c *ApiService) FetchNetworkAccessProfileNetwork(NetworkAccessProfileSid string, Sid string) (*SupersimV1NetworkAccessProfileNetwork, error) {
	path := "/v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks/{Sid}"
	path = strings.Replace(path, "{"+"NetworkAccessProfileSid"+"}", NetworkAccessProfileSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &SupersimV1NetworkAccessProfileNetwork{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListNetworkAccessProfileNetwork'
type ListNetworkAccessProfileNetworkParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListNetworkAccessProfileNetworkParams) SetPageSize(PageSize int) *ListNetworkAccessProfileNetworkParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListNetworkAccessProfileNetworkParams) SetLimit(Limit int) *ListNetworkAccessProfileNetworkParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of NetworkAccessProfileNetwork records from the API. Request is executed immediately.
func (c *ApiService) PageNetworkAccessProfileNetwork(NetworkAccessProfileSid string, params *ListNetworkAccessProfileNetworkParams, pageToken, pageNumber string) (*ListNetworkAccessProfileNetworkResponse, error) {
	path := "/v1/NetworkAccessProfiles/{NetworkAccessProfileSid}/Networks"

	path = strings.Replace(path, "{"+"NetworkAccessProfileSid"+"}", NetworkAccessProfileSid, -1)

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

	ps := &ListNetworkAccessProfileNetworkResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists NetworkAccessProfileNetwork records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListNetworkAccessProfileNetwork(NetworkAccessProfileSid string, params *ListNetworkAccessProfileNetworkParams) ([]SupersimV1NetworkAccessProfileNetwork, error) {
	response, errors := c.StreamNetworkAccessProfileNetwork(NetworkAccessProfileSid, params)

	records := make([]SupersimV1NetworkAccessProfileNetwork, 0)
	for record := range response {
		records = append(records, record)
	}

	if err := <-errors; err != nil {
		return nil, err
	}

	return records, nil
}

// Streams NetworkAccessProfileNetwork records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamNetworkAccessProfileNetwork(NetworkAccessProfileSid string, params *ListNetworkAccessProfileNetworkParams) (chan SupersimV1NetworkAccessProfileNetwork, chan error) {
	if params == nil {
		params = &ListNetworkAccessProfileNetworkParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	recordChannel := make(chan SupersimV1NetworkAccessProfileNetwork, 1)
	errorChannel := make(chan error, 1)

	response, err := c.PageNetworkAccessProfileNetwork(NetworkAccessProfileSid, params, "", "")
	if err != nil {
		errorChannel <- err
		close(recordChannel)
		close(errorChannel)
	} else {
		go c.streamNetworkAccessProfileNetwork(response, params, recordChannel, errorChannel)
	}

	return recordChannel, errorChannel
}

func (c *ApiService) streamNetworkAccessProfileNetwork(response *ListNetworkAccessProfileNetworkResponse, params *ListNetworkAccessProfileNetworkParams, recordChannel chan SupersimV1NetworkAccessProfileNetwork, errorChannel chan error) {
	curRecord := 1

	for response != nil {
		responseRecords := response.Networks
		for item := range responseRecords {
			recordChannel <- responseRecords[item]
			curRecord += 1
			if params.Limit != nil && *params.Limit < curRecord {
				close(recordChannel)
				close(errorChannel)
				return
			}
		}

		record, err := client.GetNext(c.baseURL, response, c.getNextListNetworkAccessProfileNetworkResponse)
		if err != nil {
			errorChannel <- err
			break
		} else if record == nil {
			break
		}

		response = record.(*ListNetworkAccessProfileNetworkResponse)
	}

	close(recordChannel)
	close(errorChannel)
}

func (c *ApiService) getNextListNetworkAccessProfileNetworkResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListNetworkAccessProfileNetworkResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}
