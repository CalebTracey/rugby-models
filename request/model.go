package request

type LeaderboardRequest struct {
	CompId   string `json:"compId,omitempty"`
	CompName string `json:"compName,omitempty"`
	Date     string `json:"date,omitempty"`
	Source   string `json:"source,omitempty"`
}
