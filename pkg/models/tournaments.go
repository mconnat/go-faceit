package models

type Tournaments struct {
	End   int          `json:"end"`
	Items []Tournament `json:"items"`
	Start int          `json:"start"`
}

type Tournament struct {
	AnticheatRequired           bool          `json:"anticheat_required"`
	CalculateElo                bool          `json:"calculate_elo"`
	CompetitionId               string        `json:"competition_id"`
	CoverImage                  string        `json:"cover_image"`
	Custom                      bool          `json:"custom"`
	Description                 string        `json:"description"`
	FaceITUrl                   string        `json:"faceit_url"`
	FeaturedImage               string        `json:"featured_image"`
	GameData                    Game          `json:"game_data"`
	GameId                      string        `json:"game_id"`
	InviteType                  string        `json:"invite_type"`
	MatchType                   string        `json:"match_type"`
	MaxSkill                    int           `json:"max_skill"`
	MembershipType              string        `json:"membership_type"`
	MinSkill                    int           `json:"min_skill"`
	Name                        string        `json:"name"`
	NumberOfPlayers             int           `json:"number_of_players"`
	NumberOfPlayersCheckedin    int           `json:"number_of_players_checkedin"`
	NumberOfPlayersJoined       int           `json:"number_of_players_joined"`
	NumberOfPlayersParticipants int           `json:"number_of_players_participants"`
	OrganizerData               Organizer     `json:"organizer_data"`
	OrganizerId                 string        `json:"organizer_id"`
	PrizeType                   string        `json:"prize_type"`
	Region                      string        `json:"region"`
	Rounds                      []interface{} `json:"rounds"`
	Rule                        string        `json:"rule"`
	StartedAt                   int           `json:"started_at"`
	Status                      string        `json:"status"`
	SubstitutesAllowed          int           `json:"substitutes_allowed"`
	SubstitutionsAllowed        int           `json:"substitutions_allowed"`
	SubscriptionsCount          int           `json:"subscriptions_count"`
	TeamSize                    int           `json:"team_size"`
	TournamentId                string        `json:"tournament_id"`
	WhitelistCountries          []string      `json:"whitelist_countries"`
}

type TournamentBrackets struct {
	Game    string  `json:"game"`
	Matches []Match `json:"matches"`
	Name    string  `json:"name"`
	Rounds  []struct {
		BestOf               int    `json:"best_of"`
		Label                string `json:"label"`
		Matches              int    `json:"matches"`
		Round                int    `json:"round"`
		StartTime            int    `json:"start_time"`
		StartsAsap           bool   `json:"starts_asap"`
		SubstitutionTime     int    `json:"substitution_time"`
		SubstitutionsAllowed bool   `json:"substitutions_allowed"`
	} `json:"rounds"`
	Status string `json:"status"`
}

type TournamentTeams struct {
	CheckedIn []struct {
		Nickname   string `json:"nickname"`
		SkillLevel int    `json:"skill_level"`
		SubsDone   int    `json:"subs_done"`
		TeamId     string `json:"team_id"`
		TeamLeader string `json:"team_leader"`
		TeamType   string `json:"team_type"`
	} `json:"checked_in"`
	Finished []struct {
		Nickname   string `json:"nickname"`
		SkillLevel int    `json:"skill_level"`
		SubsDone   int    `json:"subs_done"`
		TeamId     string `json:"team_id"`
		TeamLeader string `json:"team_leader"`
		TeamType   string `json:"team_type"`
	} `json:"finished"`
	Joined []struct {
		Nickname   string `json:"nickname"`
		SkillLevel int    `json:"skill_level"`
		SubsDone   int    `json:"subs_done"`
		TeamId     string `json:"team_id"`
		TeamLeader string `json:"team_leader"`
		TeamType   string `json:"team_type"`
	} `json:"joined"`
	Started []struct {
		Nickname   string `json:"nickname"`
		SkillLevel int    `json:"skill_level"`
		SubsDone   int    `json:"subs_done"`
		TeamId     string `json:"team_id"`
		TeamLeader string `json:"team_leader"`
		TeamType   string `json:"team_type"`
	} `json:"started"`
}
