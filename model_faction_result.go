/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package faceit

// # FactionResult holds detailed faction score
type FactionResult struct {
	// The score of the faction.
	Score int64 `json:"score,omitempty"`
}
