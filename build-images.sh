#!/bin/bash
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
cd project/frontend || exit
npm run build
docker build -t project-frontend .
docker tag project-frontend:latest hkotka/project-frontend
docker push hkotka/project-frontend
cd ..
cd backend || exit
docker build -t project-backend .
docker tag project-backend:latest hkotka/project-backend
docker push hkotka/project-backend
cd ..
cd daily-url-cron || exit
docker build -t daily-url-cron .
docker tag daily-url-cron:latest hkotka/daily-url-cron
docker push hkotka/daily-url-cron