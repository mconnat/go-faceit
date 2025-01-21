/*
 * Data API
 *
 * # This API provide access to FACEIT's data
 *
 * API version: 4
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package go-faceit

type Tournament struct {
	AnticheatRequired bool `json:"anticheat_required,omitempty"`
	BestOf interface{} `json:"best_of,omitempty"`
	CalculateElo bool `json:"calculate_elo,omitempty"`
	// DEPRECATED: use tournament_id instead
	CompetitionId string `json:"competition_id,omitempty"`
	CoverImage string `json:"cover_image,omitempty"`
	Custom bool `json:"custom,omitempty"`
	Description string `json:"description,omitempty"`
	FaceitUrl string `json:"faceit_url,omitempty"`
	FeaturedImage string `json:"featured_image,omitempty"`
	GameData *Game `json:"game_data,omitempty"`
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
	OrganizerData *Organizer `json:"organizer_data,omitempty"`
	OrganizerId string `json:"organizer_id,omitempty"`
	PrizeType string `json:"prize_type,omitempty"`
	Region string `json:"region,omitempty"`
	Rounds []interface{} `json:"rounds,omitempty"`
	Rule string `json:"rule,omitempty"`
	StartedAt int64 `json:"started_at,omitempty"`
	Status string `json:"status,omitempty"`
	SubstitutesAllowed int64 `json:"substitutes_allowed,omitempty"`
	SubstitutionsAllowed int64 `json:"substitutions_allowed,omitempty"`
	TeamSize int64 `json:"team_size,omitempty"`
	TotalPrize interface{} `json:"total_prize,omitempty"`
	TournamentId string `json:"tournament_id,omitempty"`
	Voting interface{} `json:"voting,omitempty"`
	WhitelistCountries []string `json:"whitelist_countries,omitempty"`
}
