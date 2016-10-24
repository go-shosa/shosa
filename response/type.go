package response

type Page struct {
	Href       string        `json:"href"`
	Offset     int           `json:"offset"`
	Limit      int           `json:"limit"`
	First      *string       `json:"first"`
	Previous   *string       `json:"previous"`
	Next       *string       `json:"next"`
	Last       *string       `json:"last"`
	Items      []interface{} `json:"items"`
	TotalItems int           `json:"total"`
}

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type Links []*Link
