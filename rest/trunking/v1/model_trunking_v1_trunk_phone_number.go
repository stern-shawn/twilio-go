/*
 * Twilio - Trunking
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
// TrunkingV1TrunkPhoneNumber struct for TrunkingV1TrunkPhoneNumber
type TrunkingV1TrunkPhoneNumber struct {
	AccountSid string `json:"AccountSid,omitempty"`
	AddressRequirements string `json:"AddressRequirements,omitempty"`
	ApiVersion string `json:"ApiVersion,omitempty"`
	Beta bool `json:"Beta,omitempty"`
	Capabilities map[string]interface{} `json:"Capabilities,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	FriendlyName string `json:"FriendlyName,omitempty"`
	Links map[string]interface{} `json:"Links,omitempty"`
	PhoneNumber string `json:"PhoneNumber,omitempty"`
	Sid string `json:"Sid,omitempty"`
	SmsApplicationSid string `json:"SmsApplicationSid,omitempty"`
	SmsFallbackMethod string `json:"SmsFallbackMethod,omitempty"`
	SmsFallbackUrl string `json:"SmsFallbackUrl,omitempty"`
	SmsMethod string `json:"SmsMethod,omitempty"`
	SmsUrl string `json:"SmsUrl,omitempty"`
	StatusCallback string `json:"StatusCallback,omitempty"`
	StatusCallbackMethod string `json:"StatusCallbackMethod,omitempty"`
	TrunkSid string `json:"TrunkSid,omitempty"`
	Url string `json:"Url,omitempty"`
	VoiceApplicationSid string `json:"VoiceApplicationSid,omitempty"`
	VoiceCallerIdLookup bool `json:"VoiceCallerIdLookup,omitempty"`
	VoiceFallbackMethod string `json:"VoiceFallbackMethod,omitempty"`
	VoiceFallbackUrl string `json:"VoiceFallbackUrl,omitempty"`
	VoiceMethod string `json:"VoiceMethod,omitempty"`
	VoiceUrl string `json:"VoiceUrl,omitempty"`
}
