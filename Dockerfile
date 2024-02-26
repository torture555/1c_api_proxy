FROM node:latest as builderNpm
LABEL author="Pavel" name="Proxy 1C API" version="0.1.0"
WORKDIR /app
COPY /public/front/ ./
ENTRYPOINT ["top", "-b"]
RUN npm install
RUN npm run build

# Build Front Vue App

FROM golang:1.22
WORKDIR /app
COPY / ./
COPY --from=builderNpm /app/dist ./
RUN go build /app/cmd/app/main.go

EXPOSE 10000
EXPOSE 11021

RUN ./main