package twiliolo

import "strconv"

type RequestOption interface {
	GetValue() (string, string)
}

type Page int

func (p Page) GetValue() (string, string) {
	return "Page", strconv.Itoa(int(p))
}

type PageSize int

func (p PageSize) GetValue() (string, string) {
	return "PageSize", strconv.Itoa(int(p))
}
