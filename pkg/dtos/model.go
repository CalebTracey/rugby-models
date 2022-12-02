package dtos

type CompetitionList []Competition

type Competition struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name" validate:"required"`
}

type CompetitionLeaderboardDataList []CompetitionLeaderboardData

type CompetitionLeaderboardData struct {
	CompId   string                  `json:"compId,omitempty"`
	CompName string                  `json:"compName,omitempty"`
	Teams    TeamLeaderboardDataList `json:"teams,omitempty"`
}

type TeamLeaderboardDataList []TeamLeaderboardData

type TeamLeaderboardData struct {
	Id               string               `json:"id,omitempty"`
	Name             string               `json:"name,omitempty"`
	Abbr             string               `json:"abbr,omitempty"`
	CompetitionStats TeamCompetitionStats `json:"competitionStats"`
}

type TeamCompetitionStats struct {
	GamesPlayed       string `json:"gamesPlayed,omitempty"`
	WinCount          string `json:"winCount,omitempty"`
	DrawCount         string `json:"drawCount,omitempty"`
	LossCount         string `json:"lossCount,omitempty"`
	Bye               string `json:"bye,omitempty"`
	PointsFor         string `json:"pointsFor,omitempty"`
	PointsAgainst     string `json:"pointsAgainst,omitempty"`
	TriesFor          string `json:"triesFor,omitempty"`
	TriesAgainst      string `json:"triesAgainst,omitempty"`
	BonusPointsTry    string `json:"bonusPointsTry,omitempty"`
	BonusPointsLosing string `json:"bonusPointsLosing,omitempty"`
	BonusPoints       string `json:"bonusPoints,omitempty"`
	PointsDiff        string `json:"pointsDiff,omitempty"`
	Points            string `json:"points,omitempty"`
}
