/*
 * Twilio - Accounts
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
// AccountsV1AuthTokenPromotion struct for AccountsV1AuthTokenPromotion
type AccountsV1AuthTokenPromotion struct {
	AccountSid string `json:"AccountSid,omitempty"`
	AuthToken string `json:"AuthToken,omitempty"`
	DateCreated time.Time `json:"DateCreated,omitempty"`
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	Url string `json:"Url,omitempty"`
}
