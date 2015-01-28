go build -a -tags netgo -ldflags '-s -linkmode external -extldflags -static' -o vhost-index
