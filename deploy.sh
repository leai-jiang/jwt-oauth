export GOPATH=/data/golang
export NODE_ENV=production

docker build -t sparta .

if docker ps -a | grep -i sparta; then
    docker rm -f sparta
fi

docker run -p 5005:5005 -d sparta
