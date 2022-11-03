# Go-client

Backend for web apps to demonstrate the use of blockchain. 

**This is a cheap copy of web SDK that we weren't able to work on due to our knowledge, skill and time issues**

This backend takes tx request (post) from angular web apps and broadcast them to blockchain. This backend was late decision since the intial idea was to use ts-client and make a web sdk for web apps  but it failed due to lot of current unstability. When we reached to cosmos developers, they said that they are still working on it. The other alternative was to use the CosmJS but it doesn't have proper documentation and it was too late. Moreover, we are getting an address of demo account (bob) from blockchain. In real life, web apps will have thier own user database that would have addresses of users stored on them and they will use those.

## How to run

Ensure the go.mod file matches the source code in module:
`go mod tidy`

Run:
`go run .` or `go run main.go`
