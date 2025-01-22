/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package faceit

// Promotion holds information about what is required in order for a player to be promoted to the next tier.
type Promotion struct {
	// Points needed for a player to get promoted.
	Points int64 `json:"points,omitempty"`
}
