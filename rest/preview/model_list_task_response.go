/*
 * Twilio - Preview
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListTaskResponse struct for ListTaskResponse
type ListTaskResponse struct {
	Meta ListDayResponseMeta `json:"Meta,omitempty"`
	Tasks []PreviewUnderstandAssistantTask `json:"Tasks,omitempty"`
}
