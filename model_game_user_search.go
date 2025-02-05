/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package faceit

// Here we return SkillLevel as string even if it is an int as we don't want to break the contract with devs
type GameUserSearch struct {
	Name string `json:"name,omitempty"`
	SkillLevel string `json:"skill_level,omitempty"`
}
