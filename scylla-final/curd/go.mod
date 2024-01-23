module example.com/crud

go 1.21.5

replace example.com/initdb => ../initdb

require example.com/initdb v0.0.0-00010101000000-000000000000

require (
	github.com/gocql/gocql v1.6.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
)

replace example.com/action => ../action
