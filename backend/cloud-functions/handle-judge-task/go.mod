module github.com/szpp-dev-team/szpp-judge/backend/cloud-functions/handle-judge-task

go 1.21.1

replace github.com/szpp-dev-team/szpp-judge/proto-gen/go => ../../../proto-gen/go

replace github.com/szpp-dev-team/szpp-judge/backend => ../../

require (
	github.com/GoogleCloudPlatform/functions-framework-go v1.7.4
	github.com/go-sql-driver/mysql v1.7.1
	github.com/kelseyhightower/envconfig v1.4.0
	github.com/szpp-dev-team/szpp-judge/backend v0.0.0-00010101000000-000000000000
	github.com/szpp-dev-team/szpp-judge/proto-gen/go v0.0.0-20230919095308-7923e085ba37
	google.golang.org/grpc v1.58.0
	google.golang.org/protobuf v1.31.0
)

require (
	ariga.io/atlas v0.10.2-0.20230427182402-87a07dfb83bf // indirect
	cloud.google.com/go/functions v1.13.0 // indirect
	entgo.io/ent v0.12.3 // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/cloudevents/sdk-go/v2 v2.14.0 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/hashicorp/hcl/v2 v2.13.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/samber/lo v1.38.1 // indirect
	github.com/zclconf/go-cty v1.8.0 // indirect
	go.uber.org/atomic v1.4.0 // indirect
	go.uber.org/multierr v1.1.0 // indirect
	go.uber.org/zap v1.10.0 // indirect
	golang.org/x/exp v0.0.0-20230811145659-89c5cff77bcb // indirect
	golang.org/x/mod v0.11.0 // indirect
	golang.org/x/net v0.14.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/text v0.12.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230807174057-1744710a1577 // indirect
)