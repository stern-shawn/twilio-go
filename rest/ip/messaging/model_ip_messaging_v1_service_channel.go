/*
 * Twilio - Ip_messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
import (
	"time"
)
// IpMessagingV1ServiceChannel struct for IpMessagingV1ServiceChannel
type IpMessagingV1ServiceChannel struct {
	AccountSid string `json:"AccountSid,omitempty"`
	Attributes string `json:"Attributes,omitempty"`
	CreatedBy string `json:"CreatedBy,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	Links map[string]interface{} `json:"Links,omitempty"`
	MembersCount int32 `json:"MembersCount,omitempty"`
	MessagesCount int32 `json:"MessagesCount,omitempty"`
	ServiceSid string `json:"ServiceSid,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Type string `json:"Type,omitempty"`
	UniqueName string `json:"UniqueName,omitempty"`
	Url string `json:"Url,omitempty"`
}
