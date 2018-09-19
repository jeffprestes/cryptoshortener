package cache

import (
	"github.com/go-macaron/cache"
	_ "github.com/go-macaron/cache/memcache"
	_ "github.com/go-macaron/cache/redis"
	"bitbucket.org/novatrixbr/cryptoshortner/conf"
)

var (
	// Memory Option
	Memory = cache.Options{}

	// File Option
	File   = cache.Options{
		Adapter:       "file",
		AdapterConfig: conf.Cfg.Section("").Key("cache_adpter_config").Value(),
	}

	// Redis Option
	Redis = cache.Options{
		Adapter:       "redis",
		AdapterConfig: conf.Cfg.Section("").Key("cache_adpter_config").Value(),
	}

	// Memcache Option
	Memcache = cache.Options{
		Adapter:       "memcache",
		AdapterConfig: conf.Cfg.Section("").Key("cache_adpter_config").Value(),
	}
)
