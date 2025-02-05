/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package faceit

type HistoryFaction struct {
	Avatar string `json:"avatar,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Players []MatchHistoryPlayer `json:"players,omitempty"`
	TeamId string `json:"team_id,omitempty"`
	Type_ string `json:"type,omitempty"`
}
