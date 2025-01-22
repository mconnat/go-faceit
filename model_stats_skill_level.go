/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package faceit

type StatsSkillLevel struct {
	Average int64 `json:"average,omitempty"`
	Range_ *StatsSkillLevelRange `json:"range,omitempty"`
}
