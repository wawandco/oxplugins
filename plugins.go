package plugins

import (
	"github.com/wawandco/oxpecker-plugins/lifecycle/build"
	"github.com/wawandco/oxpecker-plugins/lifecycle/dev"
	"github.com/wawandco/oxpecker-plugins/lifecycle/fix"
	"github.com/wawandco/oxpecker-plugins/lifecycle/generate"
	"github.com/wawandco/oxpecker-plugins/lifecycle/test"
	"github.com/wawandco/oxpecker-plugins/tools/ox"
	"github.com/wawandco/oxpecker-plugins/tools/packr"
	"github.com/wawandco/oxpecker-plugins/tools/pop"
	"github.com/wawandco/oxpecker-plugins/tools/pop/migrate"
	"github.com/wawandco/oxpecker-plugins/tools/refresh"
	"github.com/wawandco/oxpecker-plugins/tools/standard"
	"github.com/wawandco/oxpecker-plugins/tools/webpack"
	"github.com/wawandco/oxpecker-plugins/tools/yarn"
	"github.com/wawandco/oxpecker/plugins"
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
