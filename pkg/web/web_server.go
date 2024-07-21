package web

import (
	"github.com/labstack/echo/v4"
	"github.com/shreyashree-17/project/pkg/store"
)

func CreateWebServer(store *store.MongoStore) *Server {
	srv := new(Server)
	srv.store = store
	srv.echo = echo.New()

	srv.echo.GET(kHome, srv.Home)
	srv.echo.POST(kUserSignup, srv.UserSignup)
	srv.echo.POST(kSubscribeToBlogs, srv.SubscribeToBlogs)
	srv.echo.POST(kUnsubscribeFromBlogs, srv.UnsubscribeFromBlogs)
	srv.echo.GET(kRecentActions, srv.RecentActions)
	srv.echo.GET(kRecentActionsForUser, srv.RecentActionsForUser)
	return srv
}
