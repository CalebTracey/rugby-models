package request

import (
	"encoding/json"
	"github.com/calebtracey/rugby-models/pkg/dtos"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	"github.com/go-playground/validator/v10"
)

type LeaderboardRequest struct {
	Competitions dtos.CompetitionList `json:"competitions" validate:"required"`
	Date         string               `json:"date,omitempty"`
	Source       string               `json:"source,omitempty"`
}

func (req *LeaderboardRequest) UnmarshalJson(data []byte) *response.ErrorLog {
	*req = LeaderboardRequest{}
	if err := json.Unmarshal(data, req); err != nil {
		return &response.ErrorLog{
			StatusCode:    "400",
			RootCause:     err.Error(),
			ExceptionType: "bad request",
		}
	}
	return nil
}

func (req *LeaderboardRequest) Validate() error {
	validate := validator.New()
	return validate.Struct(req)
}
