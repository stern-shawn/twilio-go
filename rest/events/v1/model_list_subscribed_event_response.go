/*
 * Twilio - Events
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListSubscribedEventResponse struct for ListSubscribedEventResponse
type ListSubscribedEventResponse struct {
	Meta ListVersionResponseMeta `json:"Meta,omitempty"`
	Types []EventsV1SubscriptionSubscribedEvent `json:"Types,omitempty"`
}
