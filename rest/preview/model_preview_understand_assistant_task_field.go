/*
 * Twilio - Preview
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
// PreviewUnderstandAssistantTaskField struct for PreviewUnderstandAssistantTaskField
type PreviewUnderstandAssistantTaskField struct {
	AccountSid string `json:"AccountSid,omitempty"`
	AssistantSid string `json:"AssistantSid,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	FieldType string `json:"FieldType,omitempty"`
	Sid string `json:"Sid,omitempty"`
	TaskSid string `json:"TaskSid,omitempty"`
	UniqueName string `json:"UniqueName,omitempty"`
	Url string `json:"Url,omitempty"`
}
