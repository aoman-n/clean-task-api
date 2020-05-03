# Base Image
FROM node:12.14.0-slim as base
RUN apt-get update && apt-get install -y
WORKDIR /app
COPY ./front /app
RUN yarn install
EXPOSE 3000

# Develop Image
FROM base as development
ENV NODE_ENV=development
ENV PATH=/app/node_modules/.bin:$PATH
CMD ["yarn", "dev"]