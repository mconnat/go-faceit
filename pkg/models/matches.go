package models

type Matches struct {
	End   int     `json:"end"`
	Items []Match `json:"items"`
	Start int     `json:"start"`
}

type Match struct {
	BestOf                  int      `json:"best_of"`
	BroadcastStartTime      int      `json:"broadcast_start_time"`
	BroadcastStartTimeLabel string   `json:"broadcast_start_time_label"`
	CalculateElo            bool     `json:"calculate_elo"`
	ChatRoomId              string   `json:"chat_room_id"`
	CompetitionId           string   `json:"competition_id"`
	CompetitionName         string   `json:"competition_name"`
	CompetitionType         string   `json:"competition_type"`
	ConfiguredAt            int      `json:"configured_at"`
	DemoUrl                 []string `json:"demo_url"`
	FaceITUrl               string   `json:"faceit_url"`
	FinishedAt              int      `json:"finished_at"`
	Game                    string   `json:"game"`
	Group                   int      `json:"group"`
	MatchId                 string   `json:"match_id"`
	OrganizerId             string   `json:"organizer_id"`
	Region                  string   `json:"region"`
	Results                 struct {
		Score  map[string]int `json:"score"`
		Winner string         `json:"winner"`
	} `json:"results"`
	Round       int    `json:"round"`
	ScheduledAt int    `json:"scheduled_at"`
	StartedAt   int    `json:"started_at"`
	Status      string `json:"status"`
	Teams       map[string]struct {
		Avatar    string `json:"avatar"`
		FactionId string `json:"faction_id"`
		Leader    string `json:"leader"`
		Name      string `json:"name"`
		Roster    []struct {
			AnticheatRequired bool   `json:"anticheat_required"`
			Avatar            string `json:"avatar"`
			GamePlayerId      string `json:"game_player_id"`
			GamePlayerName    string `json:"game_player_name"`
			GameSkillLevel    int    `json:"game_skill_level"`
			Membership        string `json:"membership"`
			Nickname          string `json:"nickname"`
			PlayerId          string `json:"player_id"`
		} `json:"roster"`
		Substituted bool   `json:"substituted"`
		Type        string `json:"type"`
	} `json:"teams"`
	Version int `json:"version"`
	Voting  Voting `json:"voting"`
}

type Voting struct {
	Map MapStruct `json:"map"`
}

type MapStruct struct {
	Pick []string `json:"pick"`
}

type MatchStats struct {
	Rounds []struct {
		RoundStats struct {
		} `json:"round_stats"`
		Teams []struct {
			Players []struct {
				PlayerStats PlayerStats `json:"player_stats"`
			} `json:"players"`
			TeamStats TeamStats `json:"team_stats"`
		} `json:"teams"`
	} `json:"rounds"`
}
