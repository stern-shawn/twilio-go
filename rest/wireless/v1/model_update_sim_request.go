/*
 * Twilio - Wireless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// UpdateSimRequest struct for UpdateSimRequest
type UpdateSimRequest struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) to which the Sim resource should belong. The Account SID can only be that of the requesting Account or that of a [Subaccount](https://www.twilio.com/docs/iam/api/subaccounts) of the requesting Account. Only valid when the Sim resource's status is `new`. For more information, see the [Move SIMs between Subaccounts documentation](https://www.twilio.com/docs/wireless/api/sim-resource#move-sims-between-subaccounts).
	AccountSid string `json:"AccountSid,omitempty"`
	// The HTTP method we should use to call `callback_url`. Can be: `POST` or `GET`. The default is `POST`.
	CallbackMethod string `json:"CallbackMethod,omitempty"`
	// The URL we should call using the `callback_url` when the SIM has finished updating. When the SIM transitions from `new` to `ready` or from any status to `deactivated`, we call this URL when the status changes to an intermediate status (`ready` or `deactivated`) and again when the status changes to its final status (`active` or `canceled`).
	CallbackUrl string `json:"CallbackUrl,omitempty"`
	// The HTTP method we should use to call `commands_callback_url`. Can be: `POST` or `GET`. The default is `POST`.
	CommandsCallbackMethod string `json:"CommandsCallbackMethod,omitempty"`
	// The URL we should call using the `commands_callback_method` when the SIM sends a [Command](https://www.twilio.com/docs/wireless/api/command-resource). Your server should respond with an HTTP status code in the 200 range; any response body is ignored.
	CommandsCallbackUrl string `json:"CommandsCallbackUrl,omitempty"`
	// A descriptive string that you create to describe the Sim resource. It does not need to be unique.
	FriendlyName string `json:"FriendlyName,omitempty"`
	// The SID or unique name of the [RatePlan resource](https://www.twilio.com/docs/wireless/api/rateplan-resource) to which the Sim resource should be assigned.
	RatePlan string `json:"RatePlan,omitempty"`
	// Initiate a connectivity reset on the SIM. Set to `resetting` to initiate a connectivity reset on the SIM. No other value is valid.
	ResetStatus string `json:"ResetStatus,omitempty"`
	// The HTTP method we should use to call `sms_fallback_url`. Can be: `GET` or `POST`. Default is `POST`.
	SmsFallbackMethod string `json:"SmsFallbackMethod,omitempty"`
	// The URL we should call using the `sms_fallback_method` when an error occurs while retrieving or executing the TwiML requested from `sms_url`.
	SmsFallbackUrl string `json:"SmsFallbackUrl,omitempty"`
	// The HTTP method we should use to call `sms_url`. Can be: `GET` or `POST`. Default is `POST`.
	SmsMethod string `json:"SmsMethod,omitempty"`
	// The URL we should call using the `sms_method` when the SIM-connected device sends an SMS message that is not a [Command](https://www.twilio.com/docs/wireless/api/command-resource).
	SmsUrl string `json:"SmsUrl,omitempty"`
	// The new status of the Sim resource. Can be: `ready`, `active`, `suspended`, or `deactivated`.
	Status string `json:"Status,omitempty"`
	// An application-defined string that uniquely identifies the resource. It can be used in place of the `sid` in the URL path to address the resource.
	UniqueName string `json:"UniqueName,omitempty"`
	// The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`.
	VoiceFallbackMethod string `json:"VoiceFallbackMethod,omitempty"`
	// The URL we should call using the `voice_fallback_method` when an error occurs while retrieving or executing the TwiML requested from `voice_url`.
	VoiceFallbackUrl string `json:"VoiceFallbackUrl,omitempty"`
	// The HTTP method we should use when we call `voice_url`. Can be: `GET` or `POST`.
	VoiceMethod string `json:"VoiceMethod,omitempty"`
	// The URL we should call using the `voice_method` when the SIM-connected device makes a voice call.
	VoiceUrl string `json:"VoiceUrl,omitempty"`
}
