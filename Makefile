gen:
	[ -d pb ] || mkdir pb
	protoc --proto_path=proto proto/*.proto --go_out=plugins=grpc:pb
clean:
	rm -rf pb/*
run:
	go run main.go