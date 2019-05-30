"# spata" 

## Deploy

yarn build

git add . / git commit -m [commit-message] git push 

jenkins shell

``
#!/bin/sh
DEPLOY_PATH=/data/golang/src/sparta

rm -rf $DEPLOY_PATH/*

cp -r * $DEPLOY_PATH

cd $DEPLOY_PATH

sh deploy.sh
``
