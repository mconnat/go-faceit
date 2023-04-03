package models

type Hubs struct {
	End   int   `json:"end"`
	Items []Hub `json:"items"`
	Start int   `json:"start"`
}

type Hub struct {
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	ChatRoomId      string `json:"chat_room_id"`
	CoverImage      string `json:"cover_image"`
	Description     string `json:"description"`
	FaceITUrl       string `json:"faceit_url"`
	GameData        struct {
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
	GameId         string `json:"game_id"`
	HubId          string `json:"hub_id"`
	JoinPermission string `json:"join_permission"`
	MaxSkillLevel  int    `json:"max_skill_level"`
	MinSkillLevel  int    `json:"min_skill_level"`
	Name           string `json:"name"`
	OrganizerData  struct {
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
	OrganizerId   string `json:"organizer_id"`
	PlayersJoined int    `json:"players_joined"`
	Region        string `json:"region"`
	RuleId        string `json:"rule_id"`
}

type HubMembers struct {
	End   int `json:"end"`
	Items []struct {
		Avatar    string   `json:"avatar"`
		FaceITUrl string   `json:"faceit_url"`
		Nickname  string   `json:"nickname"`
		Roles     []string `json:"roles"`
		UserId    string   `json:"user_id"`
	} `json:"items"`
	Start int `json:"start"`
}

type HubRoles struct {
	End   int `json:"end"`
	Items []struct {
		Color         string `json:"color"`
		Name          string `json:"name"`
		Ranking       int    `json:"ranking"`
		RoleId        string `json:"role_id"`
		VisibleOnChat bool   `json:"visible_on_chat"`
	} `json:"items"`
	Start int `json:"start"`
}

type HubRules struct {
	Body      string `json:"body"`
	Game      string `json:"game"`
	Name      string `json:"name"`
	Organizer string `json:"organizer"`
	RuleId    string `json:"rule_id"`
}

type HubStats struct {
	GameId  string `json:"game_id"`
	Players []struct {
		Nickname string      `json:"nickname"`
		PlayerId string      `json:"player_id"`
		Stats    PlayerStats `json:"stats"`
	} `json:"players"`
}
