package response

import "github.com/calebtracey/rugby-models/models"

type LeaderboardResponse struct {
	Id      string                         `json:"id,omitempty"`
	Name    string                         `json:"name,omitempty"`
	Teams   models.TeamLeaderboardDataList `json:"teams,omitempty"`
	Message Message                        `json:"message,omitempty"`
}

type Message struct {
	ErrorLog  ErrorLogs `json:"errorLog,omitempty"`
	HostName  string    `json:"hostName,omitempty"`
	Status    string    `json:"status,omitempty"`
	TimeTaken string    `json:"timeTaken,omitempty"`
	Count     int       `json:"count,omitempty"`
}

type ErrorLogs []ErrorLog

type ErrorLog struct {
	Scope      string `json:"scope,omitempty"`
	StatusCode string `json:"status,omitempty"`
	Trace      string `json:"trace,omitempty"`
	RootCause  string `json:"rootCause,omitempty"`
	Query      string `json:"query,omitempty"`
}
