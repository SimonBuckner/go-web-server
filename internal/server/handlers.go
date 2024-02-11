package server

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
	"github.com/simonbuckner/go-web-server/ui"
)

type Pages struct {
	ui  *ui.UI
	log echo.Logger
}

func RegisterPageHandlers(ui *ui.UI, l echo.Logger) *Pages {
	out := Pages{
		ui:  ui,
		log: l,
	}
	return &out
}

func (p *Pages) Render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Writer.WriteHeader(statusCode)
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}

func (p *Pages) Home(c echo.Context) error {
	p.log.Info("Home page")
	return p.Render(c, http.StatusOK, p.ui.Pages.Home())
}
