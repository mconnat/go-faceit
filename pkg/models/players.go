package models

type Players struct {
	End   int      `json:"end"`
	Items []Player `json:"items"`
	Start int      `json:"start"`
}

type Player struct {
	Avatar             string   `json:"avatar"`
	Country            string   `json:"country"`
	CoverFeaturedImage string   `json:"cover_featured_image"`
	CoverImage         string   `json:"cover_image"`
	FaceITUrl          string   `json:"faceit_url"`
	FriendsIds         []string `json:"friends_ids"`
	Games              map[string]struct {
		FaceitElo       int    `json:"faceit_elo"`
		GamePlayerId    string `json:"game_player_id"`
		GamePlayerName  string `json:"game_player_name"`
		GameProfileId   string `json:"game_profile_id"`
		Region          string `json:"region"`
		SkillLevel      int    `json:"skill_level"`
		SkillLevelLabel string `json:"skill_level_label"`
	} `json:"games"`
	MembershipType string            `json:"membership_type"`
	Memberships    []string          `json:"memberships"`
	NewSteamId     string            `json:"new_steam_id"`
	Nickname       string            `json:"nickname"`
	Platforms      map[string]string `json:"platforms"`
	PlayerId       string            `json:"player_id"`
	Settings       struct {
		Language string `json:"language"`
	} `json:"settings"`
	SteamId64     string `json:"steam_id_64"`
	SteamNickname string `json:"steam_nickname"`
}

type PlayerHistory struct {
	End   int `json:"end"`
	From  int `json:"from"`
	Items []struct {
		CompetitionId   string   `json:"competition_id"`
		CompetitionName string   `json:"competition_name"`
		CompetitionType string   `json:"competition_type"`
		FaceITUrl       string   `json:"faceit_url"`
		FinishedAt      int      `json:"finished_at"`
		GameId          string   `json:"game_id"`
		GameMode        string   `json:"game_mode"`
		MatchId         string   `json:"match_id"`
		MatchType       string   `json:"match_type"`
		MaxPlayers      int      `json:"max_players"`
		OrganizerId     string   `json:"organizer_id"`
		PlayingPlayers  []string `json:"playing_players"`
		Region          string   `json:"region"`
		Results         struct {
			Score  map[string]int `json:"score"`
			Winner string         `json:"winner"`
		} `json:"results"`
		StartedAt int    `json:"started_at"`
		Status    string `json:"status"`
		Teams     map[string]struct {
			Team struct {
				Avatar   string `json:"avatar"`
				Nickname string `json:"nickname"`
				Players  []struct {
					Avatar         string `json:"avatar"`
					FaceITUrl      string `json:"faceit_url"`
					GamePlayerId   string `json:"game_player_id"`
					GamePlayerName string `json:"game_player_name"`
					Nickname       string `json:"nickname"`
					PlayerId       string `json:"player_id"`
					SkillLevel     int    `json:"skill_level"`
				} `json:"players"`
				TeamId string `json:"team_id"`
				Type   string `json:"type"`
			}
		} `json:"teams"`
		TeamsSize int `json:"teams_size"`
	} `json:"items"`
	Start int `json:"start"`
	To    int `json:"to"`
}
