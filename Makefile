# generate proto
gen:
	@protoc \
		--proto_path=api/grpc \
		--go_out=internal/pkg/common/genproto --go_opt=paths=source_relative \
		--go-grpc_out=internal/pkg/common/genproto --go-grpc_opt=paths=source_relative \
		api/grpc/*.proto

# run test
test:
	@bash ./scripts/integration_test.sh

push-image:
	@bash ./scripts/push_images.sh

# Check code security using gosec
gosec:
	@gosec -fmt yaml -out gosec_output.yaml ./...