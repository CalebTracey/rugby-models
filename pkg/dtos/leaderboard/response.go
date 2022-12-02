package leaderboard

import (
	"encoding/json"
	"github.com/calebtracey/rugby-models/pkg/dtos"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	"io"
)

type ResponseList []Response

type Response struct {
	LeaderboardData dtos.CompetitionLeaderboardDataList `json:"leaderboardData"`
	Message         response.Message                    `json:"message,omitempty"`
}

func (res *Response) ResponseToJSON(w io.Writer) *response.ErrorLog {
	if err := json.NewEncoder(w).Encode(res); err != nil {
		return &response.ErrorLog{
			StatusCode:    "500",
			RootCause:     err.Error(),
			ExceptionType: "internal server error",
		}
	}
	return nil
}
