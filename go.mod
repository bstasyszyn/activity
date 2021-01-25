module github.com/go-fed/activity

go 1.14

require (
	github.com/dave/jennifer v1.3.0
	github.com/go-fed/activity/streams v0.0.0
	github.com/go-fed/httpsig v0.1.1-0.20190914113940-c2de3672e5b5
	github.com/golang/mock v1.2.0
)

replace github.com/go-fed/activity/streams => ./streams
