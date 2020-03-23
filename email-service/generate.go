package email_service

//go:generate protoc --go_out=paths=source_relative:./proto --micro_out=paths=source_relative:./proto --proto_path=../proto ../proto/email.proto
