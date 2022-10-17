FROM node:16
WORKDIR /web
COPY .npmrc package.json ./
RUN npm install
COPY . .
ARG VITE_API_BASE_URL
RUN echo VITE_API_BASE_URL=${VITE_API_BASE_URL} > .env
RUN npm run build
RUN npm install --global serve
EXPOSE 3000
CMD [ "serve", "-n", "-s", "dist" ]
