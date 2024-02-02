templ:
	@~/go/bin/templ generate
run: templ
	@go run main.go
dev:
	@~/go/bin/reflex -s -r '\.go$$' -R '^vendor/.' -R '^_.*' -R '^gen/.' -R '^views/.*\.go$$' make run
