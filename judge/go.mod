module github.com/szpp-dev-team/szpp-judge/judge

go 1.21.0

replace github.com/szpp-dev-team/szpp-judge/proto-gen/go => ../proto-gen/go

replace github.com/szpp-dev-team/szpp-judge/langs => ../langs

require (
	github.com/docker/docker v24.0.5+incompatible
	github.com/docker/go-units v0.5.0
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.5
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/stretchr/testify v1.8.4
	github.com/szpp-dev-team/szpp-judge/langs v0.0.0-00010101000000-000000000000
	github.com/szpp-dev-team/szpp-judge/proto-gen/go v0.0.0-00010101000000-000000000000
	golang.org/x/net v0.14.0
	google.golang.org/grpc v1.57.0
)

require (
	github.com/Microsoft/go-winio v0.6.1 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/docker/distribution v2.8.2+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/moby/term v0.5.0 // indirect
	github.com/morikuni/aec v1.0.0 // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.2 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	golang.org/x/mod v0.8.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/text v0.12.0 // indirect
	golang.org/x/time v0.3.0 // indirect
	golang.org/x/tools v0.6.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230807174057-1744710a1577 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gotest.tools/v3 v3.5.0 // indirect
)
