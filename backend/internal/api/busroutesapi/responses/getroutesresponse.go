package responses

type GetRoutesResponse struct {
	Route1 string `json:"route1"`
	Route2 string `json:"route2"`
	Route3 string `json:"route3"`

	NextRoute1 string `json:"nextRoute1"`
	NextRoute2 string `json:"nextRoute2"`
	NextRoute3 string `json:"nextRoute3"`
}
