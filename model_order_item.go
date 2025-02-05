/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package faceit

type OrderItem struct {
	ExternalId string `json:"external_id,omitempty"`
	ImageUrl string `json:"image_url,omitempty"`
	ItemId string `json:"item_id,omitempty"`
	Name string `json:"name,omitempty"`
	Price float32 `json:"price,omitempty"`
}
