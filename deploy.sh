export GOPATH=/data/golang
export NODE_ENV=production

docker build -t sparta .
docker run -p 5005:5005 -d sparta
