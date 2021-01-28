/*
 * Twilio - Events
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
// EventsV1Schema struct for EventsV1Schema
type EventsV1Schema struct {
	Id string `json:"Id,omitempty"`
	LastCreated time.Time `json:"LastCreated,omitempty"`
	LastVersion int32 `json:"LastVersion,omitempty"`
	Links map[string]interface{} `json:"Links,omitempty"`
	Url string `json:"Url,omitempty"`
}
