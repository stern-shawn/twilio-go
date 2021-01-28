/*
 * Twilio - Messaging
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
// MessagingV1ServicePhoneNumber struct for MessagingV1ServicePhoneNumber
type MessagingV1ServicePhoneNumber struct {
	AccountSid string `json:"AccountSid,omitempty"`
	Capabilities []string `json:"Capabilities,omitempty"`
	CountryCode string `json:"CountryCode,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	PhoneNumber string `json:"PhoneNumber,omitempty"`
	ServiceSid string `json:"ServiceSid,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Url string `json:"Url,omitempty"`
}
