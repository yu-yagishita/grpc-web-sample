# grpc-web-sample

# 使い方
## Svelte
起動
```
cd client
yarn dev
```

## gRPC + go
起動
```
cd api
go run server.go
```

grpcの起動確認をコマンドでする
```
grpc_cli ls localhost:9090 helloworld.Greeter -l
```
## Envoy
Dockerビルドから起動まで
```
cd proxy
docker build -t envoy .
docker run --name envoy -p 8080:8080 envoy

// 再起動
docker restart envoy
```

# 参考資料
「 gRPC Web 」で gRPC 実践！ Go と gRPC で WebAPI を作ってみよう！！（やっすんのエンジニア大学）  
https://youtu.be/hlyNZoaXvqU

gRPCのドキュメント  
https://grpc.io/

gRPC WebのREADME  
https://github.com/grpc/grpc-web

envoyのDockerのドキュメント  
https://www.envoyproxy.io/docs/envoy/latest/start/docker
