FROM node:16
WORKDIR /web
RUN curl -f https://get.pnpm.io/v6.16.js | node - add --global pnpm
COPY .npmrc package.json pnpm-lock.yaml ./
RUN pnpm install
COPY . .
RUN pnpm build
RUN npm install --global serve
EXPOSE 3000
CMD [ "serve", "-n", "-s", "dist" ]