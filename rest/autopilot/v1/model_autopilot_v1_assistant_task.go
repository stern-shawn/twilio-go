/*
 * Twilio - Autopilot
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
// AutopilotV1AssistantTask struct for AutopilotV1AssistantTask
type AutopilotV1AssistantTask struct {
	AccountSid string `json:"AccountSid,omitempty"`
	ActionsUrl string `json:"ActionsUrl,omitempty"`
	AssistantSid string `json:"AssistantSid,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	Links map[string]interface{} `json:"Links,omitempty"`
	Sid string `json:"Sid,omitempty"`
	UniqueName string `json:"UniqueName,omitempty"`
	Url string `json:"Url,omitempty"`
}
