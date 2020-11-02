#!/bin/zsh
cd main-app/app1 || exit
docker build -t main-app1 .
docker tag main-app1:latest hkotka/main-app1
docker push hkotka/main-app1
cd ..
cd ..
cd main-app/app2 || exit
docker build -t main-app2 .
docker tag main-app2:latest hkotka/main-app2
docker push hkotka/main-app2
cd ..
cd ..
cd ping-pong || exit
docker build -t ping-pong .
docker tag ping-pong:latest hkotka/ping-pong
docker push hkotka/ping-pong
cd ..
cd project || exit
docker build -t project .
docker tag project:latest hkotka/project
docker push hkotka/project