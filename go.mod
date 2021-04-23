module server

go 1.16

require (
github.com/aknfujii/faker v0.0.0-00010101000000-000000000000
)


replace (
github.com/aknfujii/faker => ./pkg/faker
)
