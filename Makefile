build:
	bash -c 'set -e; go run generate.go; npx hackmyresume build resume.json TO render/resume.all -t node_modules/jsonresume-theme-latte'
