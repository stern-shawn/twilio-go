/*
 * This code was generated by
 * ___ _ _ _ _ _    _ ____    ____ ____ _    ____ ____ _  _ ____ ____ ____ ___ __   __
 *  |  | | | | |    | |  | __ |  | |__| | __ | __ |___ |\ | |___ |__/ |__|  | |  | |__/
 *  |  |_|_| | |___ | |__|    |__| |  | |    |__] |___ | \| |___ |  \ |  |  | |__| |  \
 *
 * Twilio - Api
 * This is the public Twilio REST API.
 *
 * NOTE: This class is auto generated by OpenAPI Generator.
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

package openapi

// ApiV2010AvailablePhoneNumberTollFree struct for ApiV2010AvailablePhoneNumberTollFree
type ApiV2010AvailablePhoneNumberTollFree struct {
	// A formatted version of the phone number
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The phone number in E.164 format
	PhoneNumber *string `json:"phone_number,omitempty"`
	// The LATA of this phone number
	Lata *string `json:"lata,omitempty"`
	// The locality or city of this phone number's location
	Locality *string `json:"locality,omitempty"`
	// The rate center of this phone number
	RateCenter *string `json:"rate_center,omitempty"`
	// The latitude of this phone number's location
	Latitude *string `json:"latitude,omitempty"`
	// The longitude of this phone number's location
	Longitude *string `json:"longitude,omitempty"`
	// The two-letter state or province abbreviation of this phone number's location
	Region *string `json:"region,omitempty"`
	// The postal or ZIP code of this phone number's location
	PostalCode *string `json:"postal_code,omitempty"`
	// The ISO country code of this phone number
	IsoCountry *string `json:"iso_country,omitempty"`
	// The type of Address resource the phone number requires
	AddressRequirements *string `json:"address_requirements,omitempty"`
	// Whether the phone number is new to the Twilio platform
	Beta         *bool                                                                            `json:"beta,omitempty"`
	Capabilities *ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalCapabilities `json:"capabilities,omitempty"`
}
