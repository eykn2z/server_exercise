module server

go 1.16

require (
	github.com/aknfujii/faker v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
	github.com/labstack/echo v3.3.10+incompatible
	github.com/labstack/gommon v0.3.0 // indirect
	golang.org/x/crypto v0.0.0-20210421170649-83a5a9bb288b // indirect
)

replace github.com/aknfujii/faker => ./pkg/faker
