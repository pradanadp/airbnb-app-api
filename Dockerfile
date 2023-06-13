FROM golang:1.20-alpine

# create directory folder
RUN mkdir /app

# set working directory
WORKDIR /app

COPY ./ /app

RUN go mod tidy

# create executable file with name "airbnb-api"
RUN go build -o airbnb-api

# run executable file
CMD ["./airbnb-api"]