module github.com/rwinkhart/peercred-mini

go 1.24.3

require (
	github.com/Microsoft/go-winio v0.6.2
	golang.org/x/sys v0.33.0
)

replace golang.org/x/sys => github.com/rwinkhart/sys-freebsd-13-xucred v0.33.0

replace github.com/Microsoft/go-winio => github.com/rwinkhart/go-winio-easy-pipe-handles v0.0.0-20250407031321-96994a0e8410
