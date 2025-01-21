/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package go-faceit

type OrderList struct {
	End int64 `json:"end,omitempty"`
	Items []Order `json:"items,omitempty"`
	Start int64 `json:"start,omitempty"`
}
