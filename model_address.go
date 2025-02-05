/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package faceit

type Address struct {
	Address string `json:"address,omitempty"`
	City string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
	Postcode string `json:"postcode,omitempty"`
}
