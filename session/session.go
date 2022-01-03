package session

import (
	"database/sql"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
)

type Session struct {
	CookieLifetime string
	CookiePersist  string
	CookieSecure   string
	CookieName     string
	CookieDomain   string
	SessionType    string
	DBPool         *sql.DB
}

func (s *Session) InitSession() *scs.SessionManager {
	var persist, secure bool

	// how long should sessions last?
	minutes, err := strconv.Atoi(s.CookieLifetime)
	if err != nil {
		minutes = 60
	}

	// should cookies persist>
	persist = strings.ToLower(s.CookiePersist) == "true"

	// must cookies be secure?
	secure = strings.ToLower(s.CookieSecure) == "true"

	// create session
	session := scs.New()
	session.Lifetime = time.Duration(minutes) * time.Minute
	session.Cookie.Persist = persist
	session.Cookie.Name = s.CookieName
	session.Cookie.Secure = secure
	session.Cookie.Domain = s.CookieDomain
	session.Cookie.SameSite = http.SameSiteLaxMode

	// which session store?
	switch strings.ToLower(s.SessionType) {
	case "redis":

	case "mysql", "mariadb":
		session.Store = mysqlstore.New(s.DBPool)

	case "postgres", "postgresql":
		session.Store = postgresstore.New(s.DBPool)

	default:
		// cookie
	}

	return session
}
