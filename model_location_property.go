/*
 * Messages API
 *
 * The Messaging API is a new API that consolidates all messaging channels. It encapsulates a user (developer) from having to use multiple APIs to interact with our various channels such as SMS, MMS, Viber, Facebook Messenger, etc. The API normalises information across all channels to abstracted to, from and content. This API is currently Beta.
 *
 * API version: 0.3.0
 * Contact: devrel@nexmo.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type LocationProperty struct {
	// The latitude of the location attachment.
	Lat string `json:"lat,omitempty"`
	// The longitude of the location attachment.
	Long string `json:"long,omitempty"`
	// Depending on the provider, this can either be the location on a map or the website of the business at this location.
	Url string `json:"url,omitempty"`
	// The address of the location attachment.
	Address string `json:"address,omitempty"`
	// The name of the location attachment.
	Name string `json:"name,omitempty"`
}
