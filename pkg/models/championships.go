package models

type Championships struct {
	End   int            `json:"end"`
	Items []Championship `json:"items"`
	Start int            `json:"start"`
}

type Championship struct {
	AnticheatRequired    bool   `json:"anticheat_required"`
	Avatar               string `json:"avatar"`
	BackgroundImage      string `json:"background_image"`
	ChampionshipId       string `json:"championship_id"`
	ChampionshipStart    int    `json:"championship_start"`
	CheckinClear         int    `json:"checkin_clear"`
	CheckinEnabled       bool   `json:"checkin_enabled"`
	CheckinStart         int    `json:"checkin_start"`
	CoverImage           string `json:"cover_image"`
	CurrentSubscriptions int    `json:"current_subscriptions"`
	Description          string `json:"description"`
	FaceITUrl            string `json:"faceit_url"`
	Featured             bool   `json:"featured"`
	Full                 bool   `json:"full"`
	GameData             struct {
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
	} `json:"game_data"`
	GameId     string `json:"game_id"`
	Id         string `json:"id"`
	JoinChecks struct {
		AllowedTeamTypes                []string `json:"allowed_team_types"`
		BlacklistGeoCountries           []string `json:"blacklist_geo_countries"`
		JoinPolicy                      string   `json:"join_policy"`
		MaxSkillLevel                   int      `json:"max_skill_level"`
		MembershipType                  string   `json:"membership_type"`
		MinSkillLevel                   int      `json:"min_skill_level"`
		WhitelistGeoCountries           []string `json:"whitelist_geo_countries"`
		WhitelistGeoCountriesMinPlayers int      `json:"whitelist_geo_countries_min_players"`
	} `json:"join_checks"`
	Name          string `json:"name"`
	OrganizerData struct {
		Avatar         string `json:"avatar"`
		Cover          string `json:"cover"`
		Description    string `json:"description"`
		Facebook       string `json:"facebook"`
		FaceITUrl      string `json:"faceit_url"`
		FollowersCount int    `json:"followers_count"`
		Name           string `json:"name"`
		OrganizerId    string `json:"organizer_id"`
		Twitch         string `json:"twitch"`
		Twitter        string `json:"twitter"`
		Type           string `json:"type"`
		Vk             string `json:"vk"`
		Website        string `json:"website"`
		Youtube        string `json:"youtube"`
	} `json:"organizer_data"`
	OrganizerId string `json:"organizer_id"`
	Prizes      []struct {
		FaceitPoints int `json:"faceit_points"`
		Rank         int `json:"rank"`
	} `json:"prizes"`
	Region   string `json:"region"`
	RulesId  string `json:"rules_id"`
	Schedule map[string]struct {
		Date   int    `json:"date"`
		Status string `json:"status"`
	} `json:"schedule"`
	Screening struct {
		Enabled bool   `json:"enabled"`
		Id      string `json:"id"`
	} `json:"screening"`
	SeedingStrategy string `json:"seeding_strategy"`
	Slots           int    `json:"slots"`
	Status          string `json:"status"`
	Stream          struct {
		Active   bool   `json:"active"`
		Platform string `json:"platform"`
		Source   string `json:"source"`
		Title    string `json:"title"`
	} `json:"stream"`
	SubscriptionEnd           int  `json:"subscription_end"`
	SubscriptionStart         int  `json:"subscription_start"`
	SubscriptionsLocked       bool `json:"subscriptions_locked"`
	SubstitutionConfiguration struct {
		MaxSubstitutes   int `json:"max_substitutes"`
		MaxSubstitutions int `json:"max_substitutions"`
	} `json:"substitution_configuration"`
	TotalGroups int    `json:"total_groups"`
	TotalPrizes int    `json:"total_prizes"`
	TotalRounds int    `json:"total_rounds"`
	Type        string `json:"type"`
}

type ChampionshipSubscriptions struct {
	End   int `json:"end"`
	Items []struct {
		Coach       string   `json:"coach"`
		Coleader    string   `json:"coleader"`
		Group       int      `json:"group"`
		Leader      string   `json:"leader"`
		Roster      []string `json:"roster"`
		Status      string   `json:"status"`
		Substitutes []string `json:"substitutes"`
		Team        struct {
			Avatar      string `json:"avatar"`
			ChatRoomId  string `json:"chat_room_id"`
			CoverImage  string `json:"cover_image"`
			Description string `json:"description"`
			Facebook    string `json:"facebook"`
			FaceITUrl   string `json:"faceit_url"`
			Game        string `json:"game"`
			Leader      string `json:"leader"`
			Members     []struct {
				Avatar         string   `json:"avatar"`
				Country        string   `json:"country"`
				FaceITUrl      string   `json:"faceit_url"`
				MembershipType string   `json:"membership_type"`
				Memberships    []string `json:"memberships"`
				Nickname       string   `json:"nickname"`
				SkillLevel     int      `json:"skill_level"`
				UserId         string   `json:"user_id"`
			} `json:"members"`
			Name     string `json:"name"`
			Nickname string `json:"nickname"`
			TeamId   string `json:"team_id"`
			TeamType string `json:"team_type"`
			Twitter  string `json:"twitter"`
			Website  string `json:"website"`
			Youtube  string `json:"youtube"`
		} `json:"team"`
	} `json:"items"`
	Start int `json:"start"`
}

type ChampionshipResults struct {
	End   int `json:"end"`
	Items []struct {
		Bounds struct {
			Left  int `json:"left"`
			Right int `json:"right"`
		} `json:"bounds"`
		Placements []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
			Type string `json:"type"`
		} `json:"placements"`
	} `json:"items"`
	Start int `json:"start"`
}
