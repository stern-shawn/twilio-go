/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Conversations
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

import (
	"time"
)

// ConversationsV1ServiceUserConversation struct for ConversationsV1ServiceUserConversation
type ConversationsV1ServiceUserConversation struct {
	// The unique ID of the Account responsible for this conversation.
	AccountSid *string `json:"account_sid,omitempty"`
	// The unique ID of the Conversation Service this conversation belongs to.
	ChatServiceSid *string `json:"chat_service_sid,omitempty"`
	// The unique ID of the Conversation for this User Conversation.
	ConversationSid *string `json:"conversation_sid,omitempty"`
	// The number of unread Messages in the Conversation.
	UnreadMessagesCount *int `json:"unread_messages_count,omitempty"`
	// The index of the last read Message .
	LastReadMessageIndex *int `json:"last_read_message_index,omitempty"`
	// Participant Sid.
	ParticipantSid *string `json:"participant_sid,omitempty"`
	// The unique ID for the User.
	UserSid *string `json:"user_sid,omitempty"`
	// The human-readable name of this conversation.
	FriendlyName      *string `json:"friendly_name,omitempty"`
	ConversationState *string `json:"conversation_state,omitempty"`
	// Timer date values for this conversation.
	Timers *interface{} `json:"timers,omitempty"`
	// An optional string metadata field you can use to store any data you wish.
	Attributes *string `json:"attributes,omitempty"`
	// The date that this conversation was created.
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The date that this conversation was last updated.
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// Creator of this conversation.
	CreatedBy         *string `json:"created_by,omitempty"`
	NotificationLevel *string `json:"notification_level,omitempty"`
	// An application-defined string that uniquely identifies the Conversation resource.
	UniqueName *string `json:"unique_name,omitempty"`
	Url        *string `json:"url,omitempty"`
	// Absolute URLs to access the participant and conversation of this user conversation.
	Links *map[string]interface{} `json:"links,omitempty"`
}
