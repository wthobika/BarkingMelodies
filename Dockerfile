FROM golang:1.23.4-alpine

ENV HOME /root
WORKDIR /root

COPY go.* ./

COPY . .

RUN go mod download

RUN go build -o main main.go

EXPOSE 8080

# ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.2.1/wait /wait
# RUN chmod +x /wait

# CMD [ "/wait", "./main" ] use this once the data base is set up
CMD [ "./main" ]