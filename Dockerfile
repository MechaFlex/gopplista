# First build JS, bascially just compiling the CSS
FROM node:20 as npm
WORKDIR /app
COPY . .
RUN npm install
RUN npm run build

# Then build the Go binary
FROM golang:1.22 as build
WORKDIR /app
COPY --from=npm /app /app
ENV CGO_ENABLED=0
RUN go mod download
RUN go build
RUN mkdir /app/sqlite

# Finally, copy the binary to a new image and run it
FROM scratch
WORKDIR /app
COPY --from=build /app/gopplista /app/gopplista
COPY --from=build /app/sqlite /app/sqlite
# VOLUME /app/sqlite
EXPOSE 3333
CMD ["./gopplista"]