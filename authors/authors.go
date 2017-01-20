package api

type authors struct {
	Authors []author `json:"authors"`
}

type author struct {
	Name string `json:"name"`
}

func Authors(authorNames []string) authors {
	var a []author

	for _, name := range authorNames {
		a = append(a, author{name})
	}

	return authors{a}
}
