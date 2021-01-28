/*
 * Twilio - Autopilot
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListFieldValueResponse struct for ListFieldValueResponse
type ListFieldValueResponse struct {
	FieldValues []AutopilotV1AssistantFieldTypeFieldValue `json:"FieldValues,omitempty"`
	Meta ListAssistantResponseMeta `json:"Meta,omitempty"`
}
