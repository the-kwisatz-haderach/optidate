# syntax=docker/dockerfile:1
ARG GO_VERSION=1.21

FROM golang:${GO_VERSION} AS build

WORKDIR /app

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /optidate ./main.go

FROM golang:${GO_VERSION}-alpine

RUN addgroup -S nonroot && adduser -S user -G nonroot
USER user

COPY --from=build /app/frontend/dist ./frontend/dist
COPY --from=build /optidate ./optidate

ARG PORT
ENV PORT=${PORT}
EXPOSE ${PORT}

# set hostname to localhost
ENV HOSTNAME="0.0.0.0"


ENTRYPOINT [ "./optidate" ]

