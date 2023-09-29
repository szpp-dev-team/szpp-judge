module github.com/szpp-dev-team/szpp-judge/backend/cloud-functions/handle-judge-task

go 1.21.1

replace github.com/szpp-dev-team/szpp-judge/proto-gen/go => ../../../proto-gen/go

replace github.com/szpp-dev-team/szpp-judge/backend => ../../

require (
	github.com/GoogleCloudPlatform/functions-framework-go v1.8.0
	github.com/go-sql-driver/mysql v1.7.1
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/szpp-dev-team/szpp-judge/backend v0.0.0-00010101000000-000000000000
	github.com/szpp-dev-team/szpp-judge/proto-gen/go v0.0.0-20230925111152-01fa30ed2faf
	google.golang.org/grpc v1.58.2
	google.golang.org/protobuf v1.31.0
)

require (
	ariga.io/atlas v0.14.1 // indirect
	cloud.google.com/go v0.110.8 // indirect
	cloud.google.com/go/cloudtasks v1.12.1 // indirect
	cloud.google.com/go/compute v1.23.0 // indirect
	cloud.google.com/go/compute/metadata v0.2.3 // indirect
	cloud.google.com/go/functions v1.15.1 // indirect
	cloud.google.com/go/iam v1.1.2 // indirect
	cloud.google.com/go/storage v1.33.0 // indirect
	entgo.io/ent v0.12.4 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/cloudevents/sdk-go/v2 v2.14.0 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/golang-jwt/jwt/v5 v5.0.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/s2a-go v0.1.7 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/googleapis/enterprise-certificate-proxy v0.3.1 // indirect
	github.com/googleapis/gax-go/v2 v2.12.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.4.0 // indirect
	github.com/hashicorp/hcl/v2 v2.18.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/samber/lo v1.38.1 // indirect
	github.com/zclconf/go-cty v1.14.0 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	go.uber.org/zap v1.26.0 // indirect
	golang.org/x/crypto v0.13.0 // indirect
	golang.org/x/exp v0.0.0-20230905200255-921286631fa9 // indirect
	golang.org/x/mod v0.12.0 // indirect
	golang.org/x/net v0.15.0 // indirect
	golang.org/x/oauth2 v0.12.0 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.12.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/api v0.143.0 // indirect
	google.golang.org/appengine v1.6.8 // indirect
	google.golang.org/genproto v0.0.0-20230920204549-e6e6cdab5c13 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20230920204549-e6e6cdab5c13 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230920204549-e6e6cdab5c13 // indirect
)
