run_app:
	go run main.go

swagger:
	GO111MODULE=off swagger generate spec -o ./swagger.yml --scan-models