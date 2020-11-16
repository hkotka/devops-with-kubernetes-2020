FROM node as build
COPY rollup.config.js /svelte/
COPY package*.json /svelte/
WORKDIR /svelte
RUN ls -lha
RUN npm i

COPY ./src /svelte/src
COPY ./public /svelte/public
RUN npm run build

FROM nginx
COPY --from=build /svelte/public /usr/share/nginx/html