package authors

type authors struct {
	Authors []author `json:"authors"`
}

type author struct {
  Name string `json:"name"`
}

func Authors() authors {
  luke := author{"Luke Morton"}
  bob := author{"Bob"}
  return authors{[]author{luke, bob}}
}
