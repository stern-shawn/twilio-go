/*
 * Twilio - Messaging
 *
 * This is the public Twilio REST API.
 *
 * API version: 0.1.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// CreateBrandRegistrationsRequest struct for CreateBrandRegistrationsRequest
type CreateBrandRegistrationsRequest struct {
	// A2P Messaging Profile Bundle Sid.
	A2pProfileBundleSid string `json:"A2pProfileBundleSid"`
	// Customer Profile Bundle Sid.
	CustomerProfileBundleSid string `json:"CustomerProfileBundleSid"`
}