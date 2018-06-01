package models

type SearchDetail struct {
	Type   string
	Id     int
	Name   string
	Author string
	Url    string
}
type SearchResult struct {
	Type   string
	Id     string
	Name   string
	Author string
	Url    string
}
type DescResult struct {
	Type   string
	Id     string
	Name   string
	Author string
	Url   string
}

type Searcher interface {
	Search(q,p string)(searchResult []SearchResult)
	GetDesc(id string)(descResults DescResult)
}