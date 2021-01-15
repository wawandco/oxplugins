package migrate

import (
	"github.com/gobuffalo/packd"
	"github.com/wawandco/oxplugins/plugins"
)

func Plugins(migrations packd.Box) []plugins.Plugin {
	pl := []plugins.Plugin{
		&Command{migrations: migrations},
	}

	return pl
}
