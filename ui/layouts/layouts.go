package layouts

type Layouts struct {
	Title string
}

func New(title string) *Layouts {
	return &Layouts{
		Title: title,
	}
}
