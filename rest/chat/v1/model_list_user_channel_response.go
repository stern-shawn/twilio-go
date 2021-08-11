/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListUserChannelResponse struct for ListUserChannelResponse
type ListUserChannelResponse struct {
	Channels []ChatV1UserChannel        `json:"channels,omitempty"`
	Meta     ListCredentialResponseMeta `json:"meta,omitempty"`
}
