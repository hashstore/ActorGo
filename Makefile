gen:
	[ -d pb ] || mkdir pb
	protoc --proto_path=proto proto/*.proto --go_opt=paths=source_relative --go_out=plugins=grpc:pb
	#protoc --go_out=paths=source_relative:./pb --proto_path=proto proto/*.proto
clean:
	rm -rf pb/*
run:
	go run main.go