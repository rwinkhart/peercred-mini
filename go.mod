module github.com/rwinkhart/peercred-mini

go 1.27.3

require (
	github.com/Microsoft/go-winio v0.6.4
	golang.org/x/sys v0.49.0
)

replace golang.org/x/sys => github.com/rwinkhart/sys v0.49.0

replace github.com/Microsoft/go-winio => github.com/rwinkhart/go-winio v0.6.4
