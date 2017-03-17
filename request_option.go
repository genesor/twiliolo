package twiliolo

type RequestOption interface {
	GetValue() (string, string)
}

type Page int

func (p Page) GetValue() (string, string) {
	return "Page", string(p)
}

type PageSize int

func (p PageSize) GetValue() (string, string) {
	return "PageSize", string(p)
}
