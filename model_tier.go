/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package go-faceit

type Tier struct {
	// The name of the tier
	Name string `json:"name,omitempty"`
	// The target points for the tier
	PointsTarget int64 `json:"points_target,omitempty"`
	// The rank of the tier
	Rank int64 `json:"rank,omitempty"`
}
