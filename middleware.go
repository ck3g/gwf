package gwf

import "net/http"

func (g *GWF) SessionLoad(next http.Handler) http.Handler {
	g.InfoLog.Println("SessionLoad called")
	return g.Session.LoadAndSave(next)
}
