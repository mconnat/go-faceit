/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package faceit

type Brackets struct {
	Game string `json:"game,omitempty"`
	Matches []BracketsMatch `json:"matches,omitempty"`
	Name string `json:"name,omitempty"`
	Rounds []BracketsRound `json:"rounds,omitempty"`
	Status string `json:"status,omitempty"`
}
