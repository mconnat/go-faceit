package models

type PlayerStats map[string]interface{}

type TeamStats map[string]interface{}

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
