/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package twilio
import (
	"time"
)
// VerifyV2ServiceEntityFactor struct for VerifyV2ServiceEntityFactor
type VerifyV2ServiceEntityFactor struct {
	AccountSid string `json:"account_sid,omitempty"`
	Config map[string]interface{} `json:"config,omitempty"`
	DateCreated time.Time `json:"date_created,omitempty"`
	DateUpdated time.Time `json:"date_updated,omitempty"`
	EntitySid string `json:"entity_sid,omitempty"`
	FactorType string `json:"factor_type,omitempty"`
	FriendlyName string `json:"friendly_name,omitempty"`
	Identity string `json:"identity,omitempty"`
	ServiceSid string `json:"service_sid,omitempty"`
	Sid string `json:"sid,omitempty"`
	Status string `json:"status,omitempty"`
	Url string `json:"url,omitempty"`
}