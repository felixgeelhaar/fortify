module github.com/felixgeelhaar/fortify/examples/backends/redis

go 1.24.6

require (
	github.com/felixgeelhaar/fortify v1.1.0
	github.com/felixgeelhaar/fortify/backends/redis v1.1.0
	github.com/redis/go-redis/v9 v9.7.3
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
)

// Use local modules during development
replace github.com/felixgeelhaar/fortify => ../../..

replace github.com/felixgeelhaar/fortify/backends/redis => ../../../backends/redis
