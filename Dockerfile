# FROM golang:1.19

# ##buat folder APP
# RUN mkdir /app

# ##set direktori utama
# WORKDIR /app

# ##copy seluruh file ke app
# ADD . /app

# ##buat executeable
# RUN go build -o main .

# ##jalankan executeable
# CMD ["/app/main"]

FROM golang:1.19-alpine AS builder
LABEL maintainer="Mahmuda Karima<dakasakti.id@gmail.com>"
WORKDIR /app
ADD . /app
RUN go build -o /todolist-simple

# step 2: build a small image
FROM alpine:3.16.0
WORKDIR /app
COPY --from=builder todolist-simple .
EXPOSE 3030
CMD ["./todolist-simple"]