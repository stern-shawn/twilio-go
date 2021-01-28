/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListCredentialResponse struct for ListCredentialResponse
type ListCredentialResponse struct {
	Credentials []ConversationsV1Credential `json:"Credentials,omitempty"`
	Meta ListConversationResponseMeta `json:"Meta,omitempty"`
}
