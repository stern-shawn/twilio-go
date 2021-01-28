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
// ListWebhookResponse struct for ListWebhookResponse
type ListWebhookResponse struct {
	Meta ListAssistantResponseMeta `json:"Meta,omitempty"`
	Webhooks []AutopilotV1AssistantWebhook `json:"Webhooks,omitempty"`
}
