FROM golang:1.22-alpine

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY *.go ./

# Build
RUN go build -o /go-app


ENV ADDR=8080
ENV FILE_ONE=
ENV FILE_TWO=
ENV FILE_DEST=
EXPOSE 8080

# Run
CMD /go-app  \
-source1=${FILE_ONE} \
-source2=${FILE_TWO} \
-destination=${FILE_DEST} 