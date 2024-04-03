FROM node:21-alpine

WORKDIR /app
COPY package.json .
COPY pnpm-lock.yaml .

RUN npm install -g pnpm
RUN pnpm install

COPY . .

RUN pnpm build

EXPOSE 5173

CMD ["pnpm", "preview"]