# Stage 1: Create docker image as builder (referred to later as 0) copy key files into it, then build the go app
FROM golang:1.14.9-alpine

RUN mkdir /build
ADD go.mod go.sum main.go /build/
WORKDIR /build
RUN go build

# We now have a go executable in the /build folder

# Stage 2: Create new docker image
FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser

# Now we copy that go executable into our new image /app folder 
COPY --from=0 /build/tufferina /app/

# Now we copy the templates folder too
COPY templates/ /app/templates

WORKDIR /app

ENV PORT 80
EXPOSE 80

CMD ["./tufferina"]