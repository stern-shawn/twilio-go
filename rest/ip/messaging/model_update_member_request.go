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
// UpdateMemberRequest struct for UpdateMemberRequest
type UpdateMemberRequest struct {
	Attributes string `json:"Attributes,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	LastConsumedMessageIndex *int32 `json:"LastConsumedMessageIndex,omitempty"`
	LastConsumptionTimestamp time.Time `json:"LastConsumptionTimestamp,omitempty"`
	RoleSid string `json:"RoleSid,omitempty"`
}
