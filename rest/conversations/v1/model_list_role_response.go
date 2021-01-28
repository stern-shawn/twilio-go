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
// ListRoleResponse struct for ListRoleResponse
type ListRoleResponse struct {
	Meta ListConversationResponseMeta `json:"Meta,omitempty"`
	Roles []ConversationsV1Role `json:"Roles,omitempty"`
}
