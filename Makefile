build:
	bash -c 'set -e; go run generate.go; npx hackmyresume build resume.json TO render/resume.all -t ../jsonresume-theme-engineering'

build-private:
	bash -c 'set -e; go run generate.go -c vars.private.json ; npx hackmyresume build resume.json TO render/resume.all -t ../jsonresume-theme-engineering'