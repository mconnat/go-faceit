package models

type Organizers struct {
	End   int       `json:"end"`
	Items Organizer `json:"items"`
	Start int       `json:"start"`
}

type Organizer struct {
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
}
