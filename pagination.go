package apigo

import "strconv"

type Query interface {
	Query(param string) string
}

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

func (o *PaginationQuerier) GetPage(queryable Query) *Page {
	page, err := strconv.ParseInt(queryable.Query("page"), 10, 32)
	if err != nil {
		return o.DefaultPage
	}
	size, err := strconv.ParseInt(queryable.Query("size"), 10, 32)
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
