/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package faceit

type ChampionshipSubscription struct {
	Coach string `json:"coach,omitempty"`
	Coleader string `json:"coleader,omitempty"`
	Group int64 `json:"group,omitempty"`
	Leader string `json:"leader,omitempty"`
	Roster []string `json:"roster,omitempty"`
	Status string `json:"status,omitempty"`
	Substitutes []string `json:"substitutes,omitempty"`
	Team *Team `json:"team,omitempty"`
}
