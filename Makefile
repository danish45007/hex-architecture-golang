gen_proto_message:
				protoc --go_out=C:\Users\DANISH\go\src\hex-arch\internal\adapters\framework\driving\grpc --proto_path=C:\Users\DANISH\go\src\hex-arch\internal\adapters\framework\driving\grpc\proto C:\Users\DANISH\go\src\hex-arch\internal\adapters\framework\driving\grpc\proto\number_msg.proto

gen_proto_service:
				protoc --go-grpc_out=require_unimplemented_servers=false:C:\Users\DANISH\go\src\hex-arch\internal\adapters\framework\driving\grpc --proto_path=C:\Users\DANISH\go\src\hex-arch\internal\adapters\framework\driving\grpc\proto C:\Users\DANISH\go\src\hex-arch\internal\adapters\framework\driving\grpc\proto\arithmetic_svc.proto