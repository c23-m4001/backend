FROM  golang:1.17-alpine as builder

WORKDIR /project/capstone

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .
RUN go build -tags http -o /project/capstone/build/capstone .

# to fix timezone not loaded
RUN apk add --no-cache tzdata

FROM alpine:latest
COPY --from=builder /project/capstone/build/capstone /project/capstone/build/capstone

EXPOSE 8080
ENTRYPOINT [ "cd" "/project/capstone/build", "&&" , "./capstone", "http" ]
