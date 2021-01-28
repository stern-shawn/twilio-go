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
import (
	"time"
)
// UpdateServiceConversationParticipantRequest struct for UpdateServiceConversationParticipantRequest
type UpdateServiceConversationParticipantRequest struct {
	// An optional string metadata field you can use to store any data you wish. The string value must contain structurally valid JSON if specified.  **Note** that if the attributes are not set \"{}\" will be returned.
	Attributes string `json:"Attributes,omitempty"`
	// The date that this resource was created.
	DateCreated time.Time `json:"DateCreated,omitempty"`
	// The date that this resource was last updated.
	DateUpdated time.Time `json:"DateUpdated,omitempty"`
	// A unique string identifier for the conversation participant as [Conversation User](https://www.twilio.com/docs/conversations/api/user-resource). This parameter is non-null if (and only if) the participant is using the Conversation SDK to communicate. Limited to 256 characters.
	Identity string `json:"Identity,omitempty"`
	// Index of last “read” message in the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for the Participant.
	LastReadMessageIndex *int32 `json:"LastReadMessageIndex,omitempty"`
	// Timestamp of last “read” message in the [Conversation](https://www.twilio.com/docs/conversations/api/conversation-resource) for the Participant.
	LastReadTimestamp string `json:"LastReadTimestamp,omitempty"`
	// The address of the Twilio phone number that is used in Group MMS. 'null' value will remove it.
	MessagingBindingProjectedAddress string `json:"MessagingBindingProjectedAddress,omitempty"`
	// The address of the Twilio phone number that the participant is in contact with. 'null' value will remove it.
	MessagingBindingProxyAddress string `json:"MessagingBindingProxyAddress,omitempty"`
	// The SID of a conversation-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the participant.
	RoleSid string `json:"RoleSid,omitempty"`
}
