FROM golang:1.14
WORKDIR /app
COPY . /app
COPY go.mod go.sum ./
RUN go mod download
RUN go build 
EXPOSE 80
ENTRYPOINT ./Sparkle