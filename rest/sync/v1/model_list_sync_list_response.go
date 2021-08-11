/*
 * Twilio - Sync
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListSyncListResponse struct for ListSyncListResponse
type ListSyncListResponse struct {
	Lists []SyncV1SyncList        `json:"lists,omitempty"`
	Meta  ListServiceResponseMeta `json:"meta,omitempty"`
}
