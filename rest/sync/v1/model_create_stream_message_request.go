/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// CreateStreamMessageRequest struct for CreateStreamMessageRequest
type CreateStreamMessageRequest struct {
	// A JSON string that represents an arbitrary, schema-less object that makes up the Stream Message body. Can be up to 4 KiB in length.
	Data map[string]interface{} `json:"Data"`
}
