package apigo

import (
	"net/url"
	"strconv"
)

type Page struct {
	Page int32
	Size int32
}

type PaginationQuerier struct {
	DefaultPage *Page
}

func NewPaginationQuerier(defaultPage *Page) *PaginationQuerier {
	return &PaginationQuerier{
		DefaultPage: defaultPage,
	}
}

func (o *PaginationQuerier) GetPage(qryVals url.Values) *Page {
	page, err := strconv.ParseInt(qryVals.Get("page"), 10, 32)
	if err != nil {
		return o.DefaultPage
	}
	size, err := strconv.ParseInt(qryVals.Get("size"), 10, 32)
	if err != nil {
		return o.DefaultPage
	}
	return NewPage(int32(page), int32(size))
}

func NewPage(page, size int32) *Page {
	return &Page{
		Page: page,
		Size: size,
	}
}
