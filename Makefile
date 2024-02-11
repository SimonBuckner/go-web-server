BASEDIR = ${HOME}/projects/go-web-server

help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'


confirm:
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]

## run: Runs the server ./cmd/server/
run: 
	go run ./cmd/server

## templgen: Initiates generation of TMEPL modules
templgen:
	templ generate

## air: Starts the Air hotload util running ./cmd/server/
air:
	air -c .air.toml

## get: Gets all dependancies
get: confirm
	go mod tidy

## css/bundle: connats all -part.css files into hydrate.css
css/bundle:
	./internal/css/bundlecss.sh