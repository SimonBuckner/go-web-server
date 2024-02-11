package ui

import (
	"github.com/simonbuckner/go-web-server/ui/layouts"
	"github.com/simonbuckner/go-web-server/ui/pages"
)

type UI struct {
	Layouts *layouts.Layouts
	Pages   *pages.Pages
}

func New(title string) *UI {
	layouts := layouts.New(title)
	pages := pages.New(layouts)

	return &UI{
		Layouts: layouts,
		Pages:   pages,
	}
}
