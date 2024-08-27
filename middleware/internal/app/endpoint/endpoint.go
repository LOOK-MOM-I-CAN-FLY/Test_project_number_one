package endpoint

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Service interface {
	DaysLeft() int64
}
type Endpoint struct {
	s Service
}

func New(s Service) *Endpoint {
	return &Endpoint{
		s: s,
	}
}

func (e *Endpoint) Status(ctx echo.Context) error {

	days := e.s.DaysLeft()
	server := fmt.Sprintf("Days left: %d", days)
	err := ctx.String(http.StatusOK, server)
	if err != nil {
		return err
	}

	return nil
}

