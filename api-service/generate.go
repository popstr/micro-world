package api_service

//go:generate protoc --go_out=paths=source_relative:./proto/names --micro_out=paths=source_relative:./proto/names --proto_path=../proto ../proto/names.proto
//go:generate protoc --go_out=paths=source_relative:./proto --micro_out=paths=source_relative:./proto --proto_path=../proto ../proto/email.proto
