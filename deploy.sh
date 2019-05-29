export GOPATH=/data/golang
export NODE_ENV=production

docker build -t sparta .

if docker ps | grep sparta; then
    docker rm -f sparta
fi

docker run --name sparta -p 5005:5005 -d sparta
