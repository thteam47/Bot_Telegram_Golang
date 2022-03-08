certt:
	cd cert; openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 3560 -subj "/C=VN/ST=Hai Duong/L=Hai Duong/O=THteaM/OU=THteaM/CN=*thteam.vn/emailAddress=thteam47@gmail.com" -nodes
genpb: 
	protoc --go_out=botpb --go_opt=paths=source_relative \
    --go-grpc_out=botpb --go-grpc_opt=paths=source_relative \
    proto/bot.proto
gengrpc:
	protoc -I . --grpc-gateway_out ./botpb \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    proto/bot.proto
genopenapi:
	protoc -I . --openapiv2_out ./swaggerui \
    --openapiv2_opt logtostderr=true \
    proto/bot.proto
generate:
	buf generate
BUF_VERSION:=0.43.2

run: 
	go run main.go
	
certtt:
	cd cert; ./gen.sh; cd ..

install:
	go install \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
	curl -sSL \
    	"https://github.com/bufbuild/buf/releases/download/v${BUF_VERSION}/buf-$(shell uname -s)-$(shell uname -m)" \
    	-o "$(shell go env GOPATH)/bin/buf" && \
  	chmod +x "$(shell go env GOPATH)/bin/buf"
	- If none of the above works, remove the lock files. Run in terminal:
    	sudo rm /var/lib/apt/lists/lock
		sudo rm /var/cache/apt/archives/lock
    	sudo rm /var/lib/dpkg/lock*