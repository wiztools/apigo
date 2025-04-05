package apigo

import "strconv"

type Query interface {
	Query(param string) string
}

type Page struct {
	Page int
	Size int
}

type PaginationQuerier struct {
	DefaultPage *Page
}

func NewPaginationQuerier(defaultPage *Page) *PaginationQuerier {
	return &PaginationQuerier{
		DefaultPage: defaultPage,
	}
}

func (o *PaginationQuerier) PageGet(queryable Query) *Page {
	page, err := strconv.Atoi(queryable.Query("page"))
	if err != nil {
		return o.DefaultPage
	}
	size, err := strconv.Atoi(queryable.Query("size"))
	if err != nil {
		return o.DefaultPage
	}
	return NewPage(page, size)
}

func NewPage(page, size int) *Page {
	return &Page{
		Page: page,
		Size: size,
	}
}
