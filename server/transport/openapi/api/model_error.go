/*
 * flop-strategy
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package api

type Error struct {

	// You can spot the error cause in detail by checking this code.  * `INVALID_ARGUMENT` - The arguments is invalid
	Code string `json:"code,omitempty"`

	// A developer-facing error message.
	Message string `json:"message,omitempty"`

	// A user-friendly error message.
	LocalizedMessage string `json:"localized_message,omitempty"`
}
