.PHONY: install build get

GO=go

install: build
	mv ./bin/tfs /usr/local/bin

build:
	$(GO) build -o ./bin/tfs ./main.go

get:
	$(GO) get github.com/spf13/cobra
	$(GO) get github.com/spf13/viper
	$(GO) get github.com/hashicorp/hcl/v2/hclsimple
	$(GO) get github.com/aws/aws-sdk-go-v2/config
	$(GO) get github.com/aws/aws-sdk-go-v2/service/ec2
	$(GO) get github.com/aws/aws-sdk-go/aws
	$(GO) get github.com/aws/aws-sdk-go/aws/session
	$(GO) get github.com/aws/aws-sdk-go/service/resourceexplorer2

	
