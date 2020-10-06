gen:
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:${GOPATH}/src
	#--go_opt=paths=source_relative
	#protoc --go_out=paths=source_relative:./pb --proto_path=proto proto/*.proto
clean:
	rm -f */*.pb.go
run:
	go run main.go

test:
	go test -v ./...