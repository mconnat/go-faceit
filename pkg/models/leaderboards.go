package models

type Leaderboards struct {
	End   int `json:"end"`
	Items []struct {
		CurrentStreak int `json:"current_streak"`
		Draw          int `json:"draw"`
		Lost          int `json:"lost"`
		Played        int `json:"played"`
		Player        struct {
			Avatar         string   `json:"avatar"`
			Country        string   `json:"country"`
			FaceITUrl      string   `json:"faceit_url"`
			MembershipType string   `json:"membership_type"`
			Memberships    []string `json:"memberships"`
			Nickname       string   `json:"nickname"`
			SkillLevel     int      `json:"skill_level"`
			UserId         string   `json:"user_id"`
		} `json:"player"`
		Points   int     `json:"points"`
		Position int     `json:"position"`
		WinRate  float64 `json:"win_rate"`
		Won      int     `json:"won"`
	} `json:"items"`
	Leaderboard struct {
		CompetitionId   string `json:"competition_id"`
		CompetitionType string `json:"competition_type"`
		EndDate         int    `json:"end_date"`
		GameId          string `json:"game_id"`
		Group           int    `json:"group"`
		LeaderboardId   string `json:"leaderboard_id"`
		LeaderboardMode string `json:"leaderboard_mode"`
		LeaderboardName string `json:"leaderboard_name"`
		LeaderboardType string `json:"leaderboard_type"`
		MinMatches      int    `json:"min_matches"`
		PointsPerDraw   int    `json:"points_per_draw"`
		PointsPerLoss   int    `json:"points_per_loss"`
		PointsPerWin    int    `json:"points_per_win"`
		PointsType      string `json:"points_type"`
		RankingBoost    int    `json:"ranking_boost"`
		RankingType     string `json:"ranking_type"`
		Region          string `json:"region"`
		Round           int    `json:"round"`
		Season          int    `json:"season"`
		StartDate       int    `json:"start_date"`
		StartingPoints  int    `json:"starting_points"`
		Status          string `json:"status"`
	} `json:"leaderboard"`
	Start int `json:"start"`
}

type GroupLeaderboards struct {
	End   int `json:"end"`
	Items []struct {
		CompetitionId   string `json:"competition_id"`
		CompetitionType string `json:"competition_type"`
		EndDate         int    `json:"end_date"`
		GameId          string `json:"game_id"`
		Group           int    `json:"group"`
		LeaderboardId   string `json:"leaderboard_id"`
		LeaderboardMode string `json:"leaderboard_mode"`
		LeaderboardName string `json:"leaderboard_name"`
		LeaderboardType string `json:"leaderboard_type"`
		MinMatches      int    `json:"min_matches"`
		PointsPerDraw   int    `json:"points_per_draw"`
		PointsPerLoss   int    `json:"points_per_loss"`
		PointsPerWin    int    `json:"points_per_win"`
		PointsType      string `json:"points_type"`
		RankingBoost    int    `json:"ranking_boost"`
		RankingType     string `json:"ranking_type"`
		Region          string `json:"region"`
		Round           int    `json:"round"`
		Season          int    `json:"season"`
		StartDate       int    `json:"start_date"`
		StartingPoints  int    `json:"starting_points"`
		Status          string `json:"status"`
	} `json:"items"`
	Start int `json:"start"`
}
