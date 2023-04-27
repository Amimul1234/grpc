gen:
	protoc --proto_path=proto --go_out=pb --go-grpc_out=pb proto/*.proto

clean:
	rm pb/*.go


run:
	go run main.go