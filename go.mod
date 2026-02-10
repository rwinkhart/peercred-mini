module github.com/rwinkhart/peercred-mini

go 1.25.7

require (
	github.com/Microsoft/go-winio v0.6.2
	golang.org/x/sys v0.41.0
)

replace golang.org/x/sys => github.com/rwinkhart/sys v0.41.0

replace github.com/Microsoft/go-winio => github.com/rwinkhart/go-winio v0.1.1
