package server

type Controller interface {
	Routes() (routes []Route)
}
