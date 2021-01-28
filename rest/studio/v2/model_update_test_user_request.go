/*
 * Twilio - Studio
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// UpdateTestUserRequest struct for UpdateTestUserRequest
type UpdateTestUserRequest struct {
	// List of test user identities that can test draft versions of the flow.
	TestUsers []string `json:"TestUsers"`
}
