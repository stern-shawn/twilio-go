/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListFactorResponse struct for ListFactorResponse
type ListFactorResponse struct {
	Factors []VerifyV2Factor                    `json:"factors,omitempty"`
	Meta    ListVerificationAttemptResponseMeta `json:"meta,omitempty"`
}
