package pages

import "github.com/simonbuckner/go-web-server/ui/layouts"

type Pages struct {
	layouts *layouts.Layouts
}

func New(layouts *layouts.Layouts) *Pages {
	return &Pages{
		layouts: layouts,
	}

}
