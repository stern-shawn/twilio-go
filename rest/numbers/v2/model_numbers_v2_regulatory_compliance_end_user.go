/*
 * Twilio - Numbers
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
// NumbersV2RegulatoryComplianceEndUser struct for NumbersV2RegulatoryComplianceEndUser
type NumbersV2RegulatoryComplianceEndUser struct {
	AccountSid string `json:"AccountSid,omitempty"`
	Attributes map[string]interface{} `json:"Attributes,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Type string `json:"Type,omitempty"`
	Url string `json:"Url,omitempty"`
}
