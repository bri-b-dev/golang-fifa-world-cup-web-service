# build-stage
FROM golang:alpine AS build

WORKDIR /src

COPY . .

# force rebuilding swagger in case of changes
RUN go get github.com/swaggo/swag/cmd/swag
RUN swag init

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/webservice

# run-stage
FROM scratch
COPY --from=build /src/bin/webservice ./webservice
COPY --from=build /src/winners.json ./winners.json
ENTRYPOINT [ "./webservice" ]