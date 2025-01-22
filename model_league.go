/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package faceit

// The League holds league information.
type League struct {
	// The divisions of the league.
	Divisions []Division `json:"divisions,omitempty"`
	// The game of the league.
	Game string `json:"game,omitempty"`
	// The id of the league.
	Id string `json:"id,omitempty"`
	// The minimum matches of the league.
	MinMatches int64 `json:"min_matches,omitempty"`
	// The region of the league.
	Region string `json:"region,omitempty"`
	Season *Season `json:"season,omitempty"`
}
