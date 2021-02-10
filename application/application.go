package application

import (
	"github.com/gin-gonic/gin"
	"github.com/kainonly/gin-extra/authx"
	"github.com/kainonly/gin-extra/mvcx"
	"github.com/kainonly/gin-extra/tokenx"
	"github.com/kainonly/gin-extra/typ"
	"lab-api/application/common"
	"lab-api/application/controller"
	"lab-api/application/controller/acl"
	"lab-api/application/controller/policy"
	"lab-api/application/controller/resource"
	"lab-api/application/controller/role"
	"lab-api/routes"
	"net/http"
)

func Application(router *gin.Engine, dependency common.Dependency) {
	cfg := dependency.Config
	tokenx.LoadKey([]byte(cfg.App.Key))
	router.GET("/", routes.Default)
	system := router.Group("/system")
	{
		auth := authx.AuthVerify(typ.Cookie{
			Name:     "system",
			MaxAge:   0,
			Path:     "",
			Domain:   "",
			Secure:   true,
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		}, dependency.Cache.RefreshToken)
		mvc := mvcx.Initialize(system, dependency)
		mvc.AutoController("/main", &controller.Controller{}, mvcx.Middleware{
			Handle: auth,
			Only:   []string{"Resource"},
		})
		mvc.AutoController("/acl", &acl.Controller{})
		mvc.AutoController("/resource", &resource.Controller{})
		mvc.AutoController("/policy", &policy.Controller{})
		mvc.AutoController("/role", &role.Controller{})
		mvc.AutoController("/admin", &controller.Controller{})
	}
}
