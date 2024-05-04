# Define the default target that will be run when no arguments are given to make.
all: serve

# Task to generate Go code from .templ files
generate:
	@echo "Generating Go code from .templ files..."
	templ generate

tailwind:
	@echo "Generating stylesheets..."
	tailwindcss -o assets/styles.css --minify

# Task to run the development web server
serve:
	@echo "Running the development server..."
	air

# Task to build the web server binary
build:
	@echo "Building the web server binary..."
	go build -o ./cmd/docs/webserver ./cmd/docs/main.go

.PHONY: generate serve build