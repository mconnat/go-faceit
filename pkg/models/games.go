package models

type Games struct {
	End   int    `json:"end"`
	Items []Game `json:"items"`
	Start int    `json:"start"`
}

type Game struct {
	Assets struct {
		Cover        string `json:"cover"`
		FeaturedImgL string `json:"featured_img_l"`
		FeaturedImgM string `json:"featured_img_m"`
		FeaturedImgS string `json:"featured_img_s"`
		FlagImgIcon  string `json:"flag_img_icon"`
		FlagImgL     string `json:"flag_img_l"`
		FlagImgM     string `json:"flag_img_m"`
		FlagImgS     string `json:"flag_img_s"`
		LandingPage  string `json:"landing_page"`
	} `json:"assets"`
	GameId       string   `json:"game_id"`
	LongLabel    string   `json:"long_label"`
	Order        int      `json:"order"`
	ParentGameId string   `json:"parent_game_id"`
	Platforms    []string `json:"platforms"`
	Regions      []string `json:"regions"`
	ShortLabel   string   `json:"short_label"`
}
