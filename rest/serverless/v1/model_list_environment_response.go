/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListEnvironmentResponse struct for ListEnvironmentResponse
type ListEnvironmentResponse struct {
	Environments []ServerlessV1ServiceEnvironment `json:"Environments,omitempty"`
	Meta ListServiceResponseMeta `json:"Meta,omitempty"`
}
