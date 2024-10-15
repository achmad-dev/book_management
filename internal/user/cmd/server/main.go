package main

import (
	v1 "github.com/achmad-dev/internal/user/api/route/v1"
)

// import (
// 	pb "github.com/achmad-dev/internal/pkg/common/genproto"
// )

func main() {
	envPath := "./.env"
	v1.ServeRoutes(envPath)
}
