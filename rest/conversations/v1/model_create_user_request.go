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
// CreateUserRequest struct for CreateUserRequest
type CreateUserRequest struct {
	// The JSON Object string that stores application-specific data. If attributes have not been set, `{}` is returned.
	Attributes string `json:"Attributes,omitempty"`
	// The string that you assigned to describe the resource.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// The application-defined string that uniquely identifies the resource's User within the [Conversation Service](https://www.twilio.com/docs/conversations/api/service-resource). This value is often a username or an email address, and is case-sensitive.
	Identity string `json:"Identity"`
	// The SID of a service-level [Role](https://www.twilio.com/docs/conversations/api/role-resource) to assign to the user.
	RoleSid string `json:"RoleSid,omitempty"`
}
