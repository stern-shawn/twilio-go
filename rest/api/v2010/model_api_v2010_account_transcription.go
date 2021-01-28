/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ApiV2010AccountTranscription struct for ApiV2010AccountTranscription
type ApiV2010AccountTranscription struct {
	AccountSid string `json:"AccountSid,omitempty"`
	ApiVersion string `json:"ApiVersion,omitempty"`
	DateCreated string `json:"DateCreated,omitempty"`
	DateUpdated string `json:"DateUpdated,omitempty"`
	Duration string `json:"Duration,omitempty"`
	Price float32 `json:"Price,omitempty"`
	PriceUnit string `json:"PriceUnit,omitempty"`
	RecordingSid string `json:"RecordingSid,omitempty"`
	Sid string `json:"Sid,omitempty"`
	Status string `json:"Status,omitempty"`
	TranscriptionText string `json:"TranscriptionText,omitempty"`
	Type string `json:"Type,omitempty"`
	Uri string `json:"Uri,omitempty"`
}
