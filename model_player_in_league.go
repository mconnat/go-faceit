/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package faceit

// PlayerInLeague holds information about a player in a league.
type PlayerInLeague struct {
	// The division name that the player is in.
	DivisionName string `json:"division_name,omitempty"`
	// The division tier that the player is in.
	DivisionTier string `json:"division_tier,omitempty"`
	// The division type that the player is in.
	DivisionType string `json:"division_type,omitempty"`
	// The leaderboard id that the player is in.
	LeaderboardId string `json:"leaderboard_id,omitempty"`
	// The points of the player in the leaderboard.
	Points int64 `json:"points,omitempty"`
	// The position of the player in the leaderboard.
	Position int64 `json:"position,omitempty"`
}
