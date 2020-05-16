# GRPC Client Example

```bash
# run
$ GRPC_HOST=localhost GRPC_PORT=50001 PORT=8000 go run main.go

# build
$ go build -o app
# start
$ GRPC_HOST=localhost GRPC_PORT=50001 PORT=8000 ./app

# build image
$ docker build -f ./docker/app/Dockerfile -t ${USERNAME}/grpc-client-example .
$ docker push ${USERNAME}/grpc-client-example

# deploy
$ cd deploy && helm upgrade --install grpc-client .
```
