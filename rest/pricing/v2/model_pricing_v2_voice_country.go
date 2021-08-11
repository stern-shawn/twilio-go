/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// PricingV2VoiceCountry struct for PricingV2VoiceCountry
type PricingV2VoiceCountry struct {
	// The name of the country
	Country *string `json:"country,omitempty"`
	// The ISO country code
	IsoCountry *string `json:"iso_country,omitempty"`
	// The absolute URL of the resource
	Url *string `json:"url,omitempty"`
}
