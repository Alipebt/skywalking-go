home="$(cd "$(dirname $0)"; pwd)"
go build ${GO_BUILD_OPTS} -o trace-nativation

./trace-nativation