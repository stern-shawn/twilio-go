/*
 * Twilio - Supersim
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
// SupersimV1Fleet struct for SupersimV1Fleet
type SupersimV1Fleet struct {
	AccountSid string `json:"AccountSid,omitempty"`
	CommandsEnabled bool `json:"CommandsEnabled,omitempty"`
	CommandsMethod string `json:"CommandsMethod,omitempty"`
	CommandsUrl string `json:"CommandsUrl,omitempty"`
	DataEnabled bool `json:"DataEnabled,omitempty"`
	DataLimit int32 `json:"DataLimit,omitempty"`
	DataMetering string `json:"DataMetering,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	NetworkAccessProfileSid string `json:"NetworkAccessProfileSid,omitempty"`
	Sid string `json:"Sid,omitempty"`
	UniqueName string `json:"UniqueName,omitempty"`
	Url string `json:"Url,omitempty"`
}
