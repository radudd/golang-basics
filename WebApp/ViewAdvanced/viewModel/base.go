package viewModel

type Base struct {
	Title string
}

func NewBase() {
	return Base{
		Title: "Lemonate Stand Supply",
	}
}