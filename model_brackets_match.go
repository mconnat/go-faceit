/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package go-faceit

type BracketsMatch struct {
	FaceitUrl string `json:"faceit_url,omitempty"`
	MatchId string `json:"match_id,omitempty"`
	Position int64 `json:"position,omitempty"`
	Results *MatchResult `json:"results,omitempty"`
	Round int64 `json:"round,omitempty"`
	State string `json:"state,omitempty"`
	Teams map[string]BracketsFaction `json:"teams,omitempty"`
}
