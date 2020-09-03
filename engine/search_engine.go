package engine

type SearchEngine interface {
	Search(query string) (*SearchResults, error)
}
