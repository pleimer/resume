Resume is rendered in a two step process. First generate the resume with the golang template `resume.json.template` using `generate.go`. Second generate visual renders with [hackmyresume](https://github.com/hacksalot/HackMyResume). 

The resume template uses the [FRESH schema definition](https://github.com/fresh-standard/fresh-resume-schema).


# Build
`make build` - builds FRESH template with public vars and renders html, pdf, png and yaml versions.


