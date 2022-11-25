package models

type TeamLeaderboardDataList []TeamLeaderboardData

type TeamLeaderboardData struct {
	Id                string `json:"id,omitempty"`
	Name              string `json:"name,omitempty"`
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

type PSQLLeaderboardDataList []PSQLLeaderboardData

type PSQLLeaderboardData struct {
	CompId            int    `json:"comp_id,omitempty"`
	CompName          string `json:"comp_name,omitempty"`
	TeamId            int    `json:"team_id,omitempty"`
	TeamName          string `json:"team_name,omitempty"`
	GamesPlayed       string `json:"games_played,omitempty"`
	WinCount          string `json:"win_count,omitempty"`
	DrawCount         string `json:"draw_count,omitempty"`
	LossCount         string `json:"loss_count,omitempty"`
	Bye               string `json:"bye,omitempty"`
	PointsFor         string `json:"points_for,omitempty"`
	PointsAgainst     string `json:"points_against,omitempty"`
	TriesFor          string `json:"tries_for,omitempty"`
	TriesAgainst      string `json:"tries_against,omitempty"`
	BonusPointsTry    string `json:"bonus_points_try,omitempty"`
	BonusPointsLosing string `json:"bonus_points_losing,omitempty"`
	BonusPoints       string `json:"bonus_points,omitempty"`
	PointsDiff        string `json:"points_diff,omitempty"`
	Points            string `json:"points,omitempty"`
}
