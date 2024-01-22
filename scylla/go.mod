module scylla

go 1.21.5

replace github.com/gocql/gocql => github.com/scylladb/gocql v1.12.0

require (
	github.com/gocql/gocql v1.6.0
	github.com/google/uuid v1.5.0
)

require (
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/go-cmp v0.5.4 // indirect
	github.com/hailocab/go-hostpool v0.0.0-20160125115350-e80d13ce29ed // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
)
