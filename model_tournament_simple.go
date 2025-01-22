/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package faceit

type TournamentSimple struct {
	AnticheatRequired bool `json:"anticheat_required,omitempty"`
	Custom bool `json:"custom,omitempty"`
	FaceitUrl string `json:"faceit_url,omitempty"`
	FeaturedImage string `json:"featured_image,omitempty"`
	GameId string `json:"game_id,omitempty"`
	InviteType string `json:"invite_type,omitempty"`
	MatchType string `json:"match_type,omitempty"`
	MaxSkill int64 `json:"max_skill,omitempty"`
	MembershipType string `json:"membership_type,omitempty"`
	MinSkill int64 `json:"min_skill,omitempty"`
	Name string `json:"name,omitempty"`
	NumberOfPlayers int64 `json:"number_of_players,omitempty"`
	NumberOfPlayersCheckedin int64 `json:"number_of_players_checkedin,omitempty"`
	NumberOfPlayersJoined int64 `json:"number_of_players_joined,omitempty"`
	NumberOfPlayersParticipants int64 `json:"number_of_players_participants,omitempty"`
	OrganizerId string `json:"organizer_id,omitempty"`
	PrizeType string `json:"prize_type,omitempty"`
	Region string `json:"region,omitempty"`
	StartedAt int64 `json:"started_at,omitempty"`
	Status string `json:"status,omitempty"`
	SubscriptionsCount int64 `json:"subscriptions_count,omitempty"`
	TeamSize int64 `json:"team_size,omitempty"`
	TotalPrize interface{} `json:"total_prize,omitempty"`
	TournamentId string `json:"tournament_id,omitempty"`
	WhitelistCountries []string `json:"whitelist_countries,omitempty"`
}
