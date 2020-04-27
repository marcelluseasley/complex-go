FROM golang:1.14-alpine 
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN apk add --no-cache bash git openssh && \
    go get github.com/sirupsen/logrus && \
    go get github.com/lib/pq && \
    go get github.com/gofiber/cors && \
    go get github.com/gofiber/fiber && \
    go get github.com/lib/pq  && \
    go build -o main .

COPY . .
CMD ["./main"]