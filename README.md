Resume is rendered in a two step process. First generate the resume with the golang template `resume.json.template` using `generate.go`. Second generate visual renders with [hackmyresume](https://github.com/hacksalot/HackMyResume). 

The resume template uses the [FRESH schema definition](https://github.com/fresh-standard/fresh-resume-schema).

This uses a custom version of the `jsonresume-theme-engineering` implementation.

Resume updates should be made to `resume.template.yaml`. 

The job does build a pdf version, however I have noticed a better PDF can be created by opening the html resume in the web browser and printing it to PDF. I suggest doing this when sending resumes to recruiters and companies.

# Build
`make build` - builds FRESH template with public vars and renders html, pdf, png and yaml versions.


