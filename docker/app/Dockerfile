# NOTE: Not working properly
FROM golang:1.14.6-alpine3.12 as build-env
WORKDIR /TheCatBreedAPI
COPY go.mod go.sum ./
RUN go mod download
COPY . .

EXPOSE 8080
CMD ["TheCatBreedAPI"]