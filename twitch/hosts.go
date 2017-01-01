package twitch

// TmiHostList is a structure representing a list of channels who are hosting others
// https://tmi.twitch.tv/hosts?include_logins=1&target=###
type TmiHostList struct {
	Hosts []TmiHost `json:"hosts,omitempty"`
}

// ChatBadge is an element of ChatBadges
// {"host_id":42142659,"target_id":12943124,"host_login":"kiyoshi_shinji","target_login":"ltzonda","host_display_name":"Kiyoshi_Shinji","target_display_name":"Ltzonda"}
type TmiHost struct {
	HostID            uint64 `json:"host_id,omitempty"`
	HostLogin         string `json:"host_login,omitempty"`
	HostDisplayName   string `json:"host_display_name,omitempty"`
	HostPartnered     bool   `json:"host_partnered,omitempty"`
	TargetID          uint64 `json:"target_id,omitempty"`
	TargetLogin       string `json:"target_login,omitempty"`
	TargetDisplayName string `json:"target_display_name,omitempty"`
}
