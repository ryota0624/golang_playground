build:
	statik -src=static && godotenv sh -c ". ./config.sh && vgo run main.go"

initialize:
	go get -u golang.org/x/vgo
	
lint:
	golint ./... && vgo tool vet ./

format:
	gofmt -s -w . 