package models

type Teams struct {
	End   int    `json:"end"`
	Items []Team `json:"items"`
	Start int    `json:"start"`
}

type Team struct {
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
	Verified string `json:"verified"`
	Website  string `json:"website"`
	Youtube  string `json:"youtube"`
}
