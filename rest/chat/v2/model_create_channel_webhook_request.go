/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateChannelWebhookRequest struct for CreateChannelWebhookRequest
type CreateChannelWebhookRequest struct {
	// The events that cause us to call the Channel Webhook. Used when `type` is `webhook`. This parameter takes only one event. To specify more than one event, repeat this parameter for each event. For the list of possible events, see [Webhook Event Triggers](https://www.twilio.com/docs/chat/webhook-events#webhook-event-trigger).
	ConfigurationFilters []string `json:"ConfigurationFilters,omitempty"`
	// The SID of the Studio [Flow](https://www.twilio.com/docs/studio/rest-api/flow) to call when an event in `configuration.filters` occurs. Used only when `type` is `studio`.
	ConfigurationFlowSid string `json:"ConfigurationFlowSid,omitempty"`
	// The HTTP method used to call `configuration.url`. Can be: `GET` or `POST` and the default is `POST`.
	ConfigurationMethod string `json:"ConfigurationMethod,omitempty"`
	// The number of times to retry the webhook if the first attempt fails. Can be an integer between 0 and 3, inclusive, and the default is 0.
	ConfigurationRetryCount int32 `json:"ConfigurationRetryCount,omitempty"`
	// A string that will cause us to call the webhook when it is present in a message body. This parameter takes only one trigger string. To specify more than one, repeat this parameter for each trigger string up to a total of 5 trigger strings. Used only when `type` = `trigger`.
	ConfigurationTriggers []string `json:"ConfigurationTriggers,omitempty"`
	// The URL of the webhook to call using the `configuration.method`.
	ConfigurationUrl string `json:"ConfigurationUrl,omitempty"`
	// The type of webhook. Can be: `webhook`, `studio`, or `trigger`.
	Type string `json:"Type"`
}
