/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListServiceConversationMessageResponse struct for ListServiceConversationMessageResponse
type ListServiceConversationMessageResponse struct {
	Messages []ConversationsV1ServiceConversationMessage `json:"messages,omitempty"`
	Meta     ListConversationResponseMeta                `json:"meta,omitempty"`
}
