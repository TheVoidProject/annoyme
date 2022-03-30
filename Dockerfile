# FROM golang:alpine AS builder

# WORKDIR /src/
# COPY main.go go.* /src/
# RUN CGO_ENABLED=0 go build -o /bin/annoyme

# FROM scratch
# COPY --from=builder /bin/annoyme /bin/annoyme
# ENTRYPOINT ["/bin/annoyme"]