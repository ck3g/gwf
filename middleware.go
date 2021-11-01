package gwf

import "net/http"

func (g *GWF) SessionLoad(next http.Handler) http.Handler {
	return g.Session.LoadAndSave(next)
}
