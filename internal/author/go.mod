module github.com/achmad-dev/internal/author

go 1.23.0

replace github.com/achmad-dev/internal/pkg => ../pkg

require (
	github.com/achmad-dev/internal/pkg v0.0.0-00010101000000-000000000000
	github.com/jmoiron/sqlx v1.4.0
	github.com/joho/godotenv v1.5.1
	github.com/sirupsen/logrus v1.9.3
	google.golang.org/grpc v1.67.1
)

require (
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/lib/pq v1.10.9 // indirect
	github.com/redis/go-redis/v9 v9.6.1 // indirect
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.17.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240814211410-ddb44dafa142 // indirect
	google.golang.org/protobuf v1.35.1 // indirect
)
