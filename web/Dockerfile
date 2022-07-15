FROM node:lts-alpine
WORKDIR /web
RUN npm install --global serve
COPY package*.json ./
RUN npm install
COPY . .
RUN npm run build
EXPOSE 3000
CMD [ "serve", "dist" ]