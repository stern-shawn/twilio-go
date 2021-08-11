/*
 * Twilio - Voice
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// VoiceV1DialingPermissionsSettings struct for VoiceV1DialingPermissionsSettings
type VoiceV1DialingPermissionsSettings struct {
	// `true` if the sub-account will inherit voice dialing permissions from the Master Project; otherwise `false`
	DialingPermissionsInheritance *bool `json:"dialing_permissions_inheritance,omitempty"`
	// The absolute URL of this resource
	Url *string `json:"url,omitempty"`
}
