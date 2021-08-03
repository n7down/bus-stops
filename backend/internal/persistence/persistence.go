package persistence

type Persistence interface {
	GetRoutes(stop int, routeTime string) *Routes
}
