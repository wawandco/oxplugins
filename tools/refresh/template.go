package refresh

var (
	// template for the YML file
	templateFile = `
app_root: .
build_target_path : ./cmd/{{.}}
ignored_folders:
- vendor
- log
- logs
- assets
- public
- grifts
- tmp
- bin
- node_modules
- .sass-cache
included_extensions:
- .go
- .env
build_path: bin
build_delay: 200ns
binary_name: tmp-build
command_flags: []
enable_colors: true
log_name: ox`
)
