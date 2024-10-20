package models

type PlayerStats map[string]interface{}

type TeamStats map[string]interface{}

type TeamStatsResponse struct {
	TeamID   string        `json:"team_id"`
	GameID   string        `json:"game_id"`
	Lifetime LifetimeStats `json:"lifetime"`
	Segments []Segment     `json:"segments"`
}

type LifetimeStats struct {
	Wins                           int           `json:"Wins,string"`
	WinRatePercent                 float64       `json:"Win Rate %,string"`
	TotalTeamUtilityDamagePerRound float64       `json:"Total Team Utility Damage/Round,string"`
	TotalTeamUtilityUsagePerRound  float64       `json:"Total Team Utility Usage/Round,string"`
	TotalTeamUtilitySuccessRate    float64       `json:"Total Team Utility Success Rate,string"`
	TeamAvgSniperKillRate          float64       `json:"Team Avg Sniper Kill Rate,string"`
	TeamAverageKRRatio             float64       `json:"Team Average K/R Ratio,string"`
	TeamAvgUtilityDamagePerRound   float64       `json:"Team Avg Utility Damage/Round,string"`
	TotalTeamEntrySuccess          float64       `json:"Total Team Entry Success,string"`
	TotalTeamEntryRate             float64       `json:"Total Team Entry Rate,string"`
	TotalHeadshotsPercent          float64       `json:"Total Headshots %,string"`
	Team1v1WinRate                 float64       `json:"Team 1v1 Win Rate,string"`
	RecentResults                  RecentResults `json:"Recent Results"`
	TotalTeam1v2WinRate            float64       `json:"Total Team 1v2 Win Rate,string"`
	TeamAvgSniperKillRatePerRound  float64       `json:"Team Avg Sniper Kill Rate per Round,string"`
	TotalTeamSniperKillRate        float64       `json:"Total Team Sniper Kill Rate,string"`
	TotalMatches                   int           `json:"Total Matches,string"`
	TeamEntrySuccess               float64       `json:"Team Entry Success,string"`
	TeamAverageHeadshotsPercent    float64       `json:"Team Average Headshots %,string"`
	TotalTeam1v1WinRate            float64       `json:"Total Team 1v1 Win Rate,string"`
	TeamAvgUtilityUsagePerRound    float64       `json:"Team Avg Utility Usage/Round,string"`
	TotalHeadshotsPerMatch         float64       `json:"Total Headshots per Match,string"`
	TotalTeamEnemiesFlashed        float64       `json:"Total Team Enemies Flashed,string"`
	TeamADR                        float64       `json:"Team ADR,string"`
	TeamAvgFlashSuccess            float64       `json:"Team Avg Flash Success,string"`
	TeamAvgEnemiesFlashed          float64       `json:"Team Avg Enemies Flashed,string"`
	TeamUtilityDamageSuccess       float64       `json:"Team Utility Damage Success,string"`
	LongestWinStreak               int           `json:"Longest Win Streak,string"`
	TeamHeadshotsPerMatch          float64       `json:"Team Headshots per Match,string"`
	CurrentWinStreak               int           `json:"Current Win Streak,string"`
	TotalTeamFlashSuccessRate      float64       `json:"Total Team Flash Success Rate,string"`
	TeamAvgUtilitySuccessPerRound  float64       `json:"Team Avg Utility Success/Round,string"`
	TeamEntryRate                  float64       `json:"Team Entry Rate,string"`
	TotalKDRatio                   float64       `json:"Total K/D Ratio,string"`
	TotalTeamFlashesPerRound       float64       `json:"Total Team Flashes/Round,string"`
	TotalKRRatio                   float64       `json:"Total K/R Ratio,string"`
	Matches                        int           `json:"Matches,string"`
	TeamAverageKDRatio             float64       `json:"Team Average K/D Ratio,string"`
	Team1v2WinRate                 float64       `json:"Team 1v2 Win Rate,string"`
	TeamAvgFlashPerRound           float64       `json:"Team Avg Flash/Round,string"`
	TotalTeamUtilityDamageSuccess  float64       `json:"Total Team Utility Damage Success,string"`
}

type RecentResults []int32

func (rr *RecentResults) UnmarshalJSON(data []byte) error {
	var results []string
	if err := json.Unmarshal(data, &results); err != nil {
		return err
	}

	for _, s := range results {
		num, err := strconv.Atoi(s)
		if err != nil {
			return fmt.Errorf("failed to convert string to int: %v", err)
		}
		*rr = append(*rr, int32(num))
	}
	return nil
}

type Segment struct {
	Type       string       `json:"type"`
	Mode       string       `json:"mode"`
	Label      string       `json:"label"`
	ImgSmall   string       `json:"img_small"`
	ImgRegular string       `json:"img_regular"`
	Stats      SegmentStats `json:"stats"`
}

type SegmentStats struct {
	WinRatePercent float64 `json:"Win Rate %,string"`
	Matches        int     `json:"Matches,string"`
	Wins           int     `json:"Wins,string"`
	TotalMatches   int     `json:"Total Matches,string"`
}

type PlayerMatchResponse struct {
	Items []PlayerMatch `json:"items"`
}

type PlayerMatch struct {
	Stats PlayerMatchStats `json:"stats"`
}

type PlayerMatchStats struct {
	CompetitionId    string `json:"Competition Id"`
	HeadshotsPercent string `json:"Headshots %"`
	KRRatio          string `json:"K/R Ratio"`
	PentaKills       string `json:"Penta Kills"`
	MatchFinishedAt  int64  `json:"Match Finished At"`
	Headshots        string `json:"Headshots"`
	PlayerId         string `json:"Player Id"`
	Winner           string `json:"Winner"`
	CreatedAt        string `json:"Created At"`
	MVPs             string `json:"MVPs"`
	Map              string `json:"Map"`
	OvertimeScore    string `json:"Overtime score"`
	FinalScore       string `json:"Final Score"`
	Rounds           string `json:"Rounds"`
	SecondHalfScore  string `json:"Second Half Score"`
	Kills            string `json:"Kills"`
	Game             string `json:"Game"`
	Deaths           string `json:"Deaths"`
	Region           string `json:"Region"`
	MatchId          string `json:"Match Id"`
	UpdatedAt        string `json:"Updated At"`
	Score            string `json:"Score"`
	Assists          string `json:"Assists"`
	Team             string `json:"Team"`
	TripleKills      string `json:"Triple Kills"`
	KDRatio          string `json:"K/D Ratio"`
	MatchRound       string `json:"Match Round"`
	BestOf           string `json:"Best Of"`
	Result           string `json:"Result"`
	Nickname         string `json:"Nickname"`
	QuadroKills      string `json:"Quadro Kills"`
	DoubleKills      string `json:"Double Kills"`
	FirstHalfScore   string `json:"First Half Score"`
	GameMode         string `json:"Game Mode"`
}
