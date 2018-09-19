package app

import (
	"bitbucket.org/novatrixbr/cryptoshortner/conf"
	"bitbucket.org/novatrixbr/cryptoshortner/handler"
	"bitbucket.org/novatrixbr/cryptoshortner/lib/auth"
	"bitbucket.org/novatrixbr/cryptoshortner/lib/cache"
	"bitbucket.org/novatrixbr/cryptoshortner/lib/contx"
	"bitbucket.org/novatrixbr/cryptoshortner/lib/cors"
	"bitbucket.org/novatrixbr/cryptoshortner/lib/template"
	mcache "github.com/go-macaron/cache"
	"github.com/go-macaron/gzip"
	"github.com/go-macaron/i18n"
	"github.com/go-macaron/jade"
	"github.com/go-macaron/session"
	"github.com/go-macaron/toolbox"
	"github.com/prometheus/client_golang/prometheus"
	"gopkg.in/macaron.v1"
)

//SetupMiddlewares configures the middlewares using in each web request
func SetupMiddlewares(app *macaron.Macaron) {
	app.Use(macaron.Logger())
	app.Use(macaron.Recovery())
	app.Use(gzip.Gziper())
	app.Use(toolbox.Toolboxer(app, toolbox.Options{
		HealthCheckers: []toolbox.HealthChecker{
			new(handler.AppChecker),
		},
	}))
	app.Use(macaron.Static("public"))
	app.Use(i18n.I18n(i18n.Options{
		Directory: "locale",
		Langs:     []string{"pt-BR", "en-US"},
		Names:     []string{"Português do Brasil", "American English"},
	}))
	app.Use(jade.Renderer(jade.Options{
		Directory: "public/templates",
		Funcs:     template.FuncMaps(),
	}))
	app.Use(macaron.Renderer(macaron.RenderOptions{
		Directory: "public/templates",
		Funcs:     template.FuncMaps(),
	}))
	//Cache in memory
	app.Use(mcache.Cacher(
		cache.Option(conf.Cfg.Section("").Key("cache_adapter").Value()),
	))
	/*
		Redis Cache
		Add this lib to import session: _ "github.com/go-macaron/cache/redis"
		Later replaces the cache in memory instructions for the lines below
		optCache := mcache.Options{
				Adapter:       conf.Cfg.Section("").Key("cache_adapter").Value(),
				AdapterConfig: conf.Cfg.Section("").Key("cache_adapter_config").Value(),
			}
		app.Use(mcache.Cacher(optCache))
	*/
	app.Use(session.Sessioner())
	app.Use(contx.Contexter())
	app.Use(cors.Cors())
	conf.LoadMongoConfig()
}

//SetupRoutes defines the routes the Web Application will respond
func SetupRoutes(app *macaron.Macaron) {
	app.Get("", handler.Index)
	app.Get("/create", handler.OpenCreatePage)
	app.Get("/:nickname", handler.GetAccount)

	//HealthChecker
	app.Get("/health", handler.HealthCheck)

	//Prometheus metrics
	app.Get("/metrics", prometheus.Handler())

	app.Group("/api", func() {
		app.Group("/v1", func() {
			app.Post("/account", handler.InsertAccount)
			app.Post("/oauth", handler.NewAppClient)
		}, auth.LoginRequiredApi)
		app.Get("/oauth", handler.Oauth)
	})

}
