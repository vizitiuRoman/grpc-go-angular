protoc \
  -I ../../../ \
  -I ${GOPATH}/src \
  --go_out="plugins=grpc:." \
  ../../../grpc-proto/**/*.proto
