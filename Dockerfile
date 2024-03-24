FROM node:latest as builderNpm
LABEL author="Pavel" name="Proxy 1C API" version="0.1.0"
WORKDIR /app
COPY /client/front/ ./
ENTRYPOINT ["top", "-b"]
RUN npm install
RUN npm run build

# Build Front Vue App

FROM golang:1.22
WORKDIR /app
COPY / ./
COPY --from=builderNpm /app/dist ./dist
RUN go build /app/cmd/app/main.go

EXPOSE 10000
EXPOSE 10001
EXPOSE 11021

CMD "./main"