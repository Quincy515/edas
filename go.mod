module edas

go 1.13

require (
	github.com/casbin/casbin v1.9.1
	github.com/casbin/casbin-server v0.0.0-20191011091846-dc83b3dae094 // indirect
	github.com/casbin/casbin/v2 v2.0.2
	github.com/go-sql-driver/mysql v1.4.1
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/jmoiron/sqlx v1.2.0
	github.com/json-iterator/go v1.1.8
	github.com/juju/ratelimit v1.0.1
	github.com/julienschmidt/httprouter v1.2.0
	github.com/micro/cli v0.2.0
	github.com/micro/go-micro v1.15.1
	github.com/micro/go-plugins v1.4.0
	go.uber.org/multierr v1.4.0 // indirect
	go.uber.org/zap v1.12.0
	golang.org/x/net v0.0.0-20191011234655-491137f69257
	golang.org/x/tools v0.0.0-20191107010934-f79515f33823 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0
	gopkg.in/yaml.v2 v2.2.5 // indirect
)

replace github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.0
