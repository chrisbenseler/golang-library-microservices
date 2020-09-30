package domain

//Router router interface
type Router interface {
}

type routerStruct struct {
}

//NewRouter create new router
func NewRouter() Router {
	return &routerStruct{}
}
