package leaderboard

import (
	"encoding/json"
	"fmt"
	"github.com/calebtracey/rugby-models/pkg/dtos"
	"github.com/calebtracey/rugby-models/pkg/dtos/response"
	"github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
	"io"
)

type Request struct {
	Competitions dtos.CompetitionList `json:"competitions" validate:"required"`
	Date         string               `json:"date,omitempty"`
	Source       string               `json:"source,omitempty"`
}

func RequestFromJSON(reader io.Reader) (request Request) {
	err := json.NewDecoder(reader).Decode(&request)
	if err != nil {
		log.Errorf("error decoding leaderboard request: %v", err)
	}
	return request
}

func (r Request) RequestToJSON(writer io.Writer) {
	err := json.NewEncoder(writer).Encode(r)
	if err != nil {
		log.Errorf("error encoding leaderboard request: %v", err)
	}
}

func (r Request) Validate() *response.ErrorLog {
	if err := validator.New().Struct(r); err != nil {
		return &response.ErrorLog{
			StatusCode:    "400",
			RootCause:     err.Error(),
			Query:         fmt.Sprintf("%s", r),
			ExceptionType: "bad request",
		}
	}
	return nil
}
