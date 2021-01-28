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
// ListSyncMapItemResponse struct for ListSyncMapItemResponse
type ListSyncMapItemResponse struct {
	Items []SyncV1ServiceSyncMapSyncMapItem `json:"Items,omitempty"`
	Meta ListServiceResponseMeta `json:"Meta,omitempty"`
}
