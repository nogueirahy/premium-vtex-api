package internal

import (
	"net/http"
)

func registerRouteGroup(mux *http.ServeMux, prefix string, handler http.Handler) *http.ServeMux {
	mux.Handle(prefix+"/", http.StripPrefix(prefix, handler))
	return mux
}

func AuthorizedRoutes(deps AppDependencies) *http.ServeMux {
	mainMux := http.NewServeMux()
	registerRouteGroup(mainMux, "/api/v1/prime", primeMux(deps))
	return mainMux
}

func primeMux(deps AppDependencies) *http.ServeMux {
	primeMux := http.NewServeMux()
	primeMux.Handle("/simulation-prime", http.HandlerFunc(deps.SimulationPrice))
	primeMux.Handle("/add-item", http.HandlerFunc(deps.AddItemPrime))
	primeMux.Handle("/remove-item", http.HandlerFunc(deps.RemoveItemPrime))

	return primeMux
}
