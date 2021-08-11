/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010Queue struct for ApiV2010Queue
type ApiV2010Queue struct {
	// The SID of the Account that created this resource
	AccountSid *string `json:"account_sid,omitempty"`
	// Average wait time of members in the queue
	AverageWaitTime *int `json:"average_wait_time,omitempty"`
	// The number of calls currently in the queue.
	CurrentSize *int `json:"current_size,omitempty"`
	// The RFC 2822 date and time in GMT that this resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that this resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// A string that you assigned to describe this resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The max number of calls allowed in the queue
	MaxSize *int `json:"max_size,omitempty"`
	// The unique string that identifies this resource
	Sid *string `json:"sid,omitempty"`
	// The URI of this resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
}
