module github.com/LinkedInLearning/hands-on-go-3080775

go 1.18

// go get [lib name]
// module tool will retrieve the module if it doesn't already have it, and cache it locally
// go mod tidy
// this will do retrieval and removal of packages as needed and keep your dependencies file clean
require (
	github.com/davecgh/go-spew v1.1.1
	github.com/jboursiquot/go-proverbs v0.0.2
	golang.org/x/exp v0.0.0-20220921164117-439092de6870
)