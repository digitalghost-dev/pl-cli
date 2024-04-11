FROM golang:1.21-alpine3.19

WORKDIR /app

ENV TERM xterm-256color
ENV COLOR_OUTPUT true

COPY . /app

RUN PATH="$PATH:~/go/bin:/usr/local/go/bin:$GOPATH/bin"

RUN go install

ENTRYPOINT ["pl-cli"]