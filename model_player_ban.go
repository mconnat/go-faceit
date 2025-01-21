/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package go-faceit

import (
	"time"
)

type PlayerBan struct {
	EndsAt time.Time `json:"ends_at,omitempty"`
	Game string `json:"game,omitempty"`
	Nickname string `json:"nickname,omitempty"`
	Reason string `json:"reason,omitempty"`
	StartsAt time.Time `json:"starts_at,omitempty"`
	Type_ string `json:"type,omitempty"`
	UserId string `json:"user_id,omitempty"`
}
