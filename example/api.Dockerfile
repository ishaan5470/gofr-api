FROM golang:1.16-alpine
WORKDIR /app
RUN apk update && apk add libc-dev && apk add gcc && apk add make
#copying our go mod and go sum into our working drectory 

COPY ./go.mod go.sum ./
RUN go mod download && go mod verify

#install compileDaemon  which allows us to wathc our gofiles
RUN go get github.com/githubnemo/CompileDaemon

COPY . .
#copying entry.sh file into the directory
COPY ./entrypoint.sh /entrypoint.sh

ADD https://raw.githubusercontent.com/eficode/wait-for/v2.1.0/wait-for /usr/local/bin/wait-for
RUN chmod +rx /usr/local/bin/wait-for /entrypoint.sh


#this commmand will run eveyrtime my dockers componse will run 

ENTRYPOINT [ "sh", "/entrypoint.sh" ]
