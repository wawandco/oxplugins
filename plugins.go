package oxplugins

import (
	"github.com/wawandco/oxpecker/plugins"
	"github.com/wawandco/oxplugins/lifecycle/build"
	"github.com/wawandco/oxplugins/lifecycle/dev"
	"github.com/wawandco/oxplugins/lifecycle/fix"
	"github.com/wawandco/oxplugins/lifecycle/generate"
	"github.com/wawandco/oxplugins/lifecycle/test"
	"github.com/wawandco/oxplugins/tools/ox"
	"github.com/wawandco/oxplugins/tools/packr"
	"github.com/wawandco/oxplugins/tools/pop"
	"github.com/wawandco/oxplugins/tools/pop/migrate"
	"github.com/wawandco/oxplugins/tools/refresh"
	"github.com/wawandco/oxplugins/tools/standard"
	"github.com/wawandco/oxplugins/tools/webpack"
	"github.com/wawandco/oxplugins/tools/yarn"
)

// All plugins in this package
var All = []plugins.Plugin{
	// Tools plugins.
	&webpack.Plugin{},
	&refresh.Plugin{},
	&packr.Plugin{},
	&pop.Plugin{},
	&migrate.Plugin{},
	&migrate.MigrateUp{},
	&migrate.MigrateDown{},
	&standard.Plugin{},
	&yarn.Plugin{},

	// Fixers
	&pop.Fixer{},
	&standard.Fixer{},

	// Generators
	&ox.Generator{},
	&ox.Builder{},

	// Developer Lifecycle plugins
	&build.Command{},
	&dev.Command{},
	&test.Command{},
	&fix.Command{},
	&generate.Command{},
}
