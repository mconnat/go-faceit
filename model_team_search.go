/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package faceit

type TeamSearch struct {
	Avatar string `json:"avatar,omitempty"`
	ChatRoomId string `json:"chat_room_id,omitempty"`
	FaceitUrl string `json:"faceit_url,omitempty"`
	Game string `json:"game,omitempty"`
	Name string `json:"name,omitempty"`
	TeamId string `json:"team_id,omitempty"`
	Verified bool `json:"verified,omitempty"`
}
