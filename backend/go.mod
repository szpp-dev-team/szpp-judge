module github.com/szpp-dev-team/szpp-judge/backend

go 1.21

replace github.com/szpp-dev-team/szpp-judge/proto-gen/go => ../proto-gen/go

require (
	cloud.google.com/go/cloudtasks v1.12.1
	cloud.google.com/go/storage v1.32.0
	connectrpc.com/connect v1.11.1
	entgo.io/ent v0.12.3
	github.com/go-sql-driver/mysql v1.7.1
	github.com/golang-jwt/jwt/v5 v5.0.0
	github.com/google/uuid v1.3.0
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/rs/cors v1.10.1
	github.com/samber/lo v1.38.1
	github.com/stretchr/testify v1.8.2
	github.com/szpp-dev-team/szpp-judge/proto-gen/go v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.13.0
	golang.org/x/exp v0.0.0-20230811145659-89c5cff77bcb
	golang.org/x/net v0.14.0
	google.golang.org/grpc v1.58.0
	google.golang.org/protobuf v1.31.0
)

require (
	ariga.io/atlas v0.10.2-0.20230427182402-87a07dfb83bf // indirect
	cloud.google.com/go v0.110.6 // indirect
	cloud.google.com/go/compute v1.23.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/iam v1.1.1 // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/s2a-go v0.1.5 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.2.5 // indirect
	github.com/googleapis/gax-go/v2 v2.12.0 // indirect
	github.com/hashicorp/hcl/v2 v2.13.0 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/zclconf/go-cty v1.8.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	golang.org/x/mod v0.11.0 // indirect
	golang.org/x/oauth2 v0.11.0 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/api v0.138.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20230803162519-f966b187b2e5 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230803162519-f966b187b2e5 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230807174057-1744710a1577 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
