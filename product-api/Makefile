check_install:
	which swagger || go install github.com/go-swagger/go-swagger/cmd/swagger@latest

swagger:
	swagger generate spec -o ./swagger.yaml --scan-models

gen_client: 
	swagger generate client -f ./../swagger.yaml -A product-api 