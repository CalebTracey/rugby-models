package leaderboard

import (
	"encoding/json"
	"github.com/calebtracey/rugby-models/pkg/dtos"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	"github.com/go-playground/validator/v10"
	"io"
)

type Request struct {
	Competitions dtos.CompetitionList `json:"competitions" validate:"required"`
	Date         string               `json:"date,omitempty"`
	Source       string               `json:"source,omitempty"`
}

func (req *Request) FromJSON(r io.Reader) *response.ErrorLog {
	e := json.NewDecoder(r)
	if err := e.Decode(req); err != nil {
		return &response.ErrorLog{
			StatusCode:    "400",
			RootCause:     err.Error(),
			ExceptionType: "bad request",
		}
	}
	return nil
}

func (req *Request) Validate() *response.ErrorLog {
	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		return &response.ErrorLog{
			StatusCode:    "400",
			RootCause:     err.Error(),
			ExceptionType: "bad request",
		}
	}
	return nil
}
