FROM cgr.dev/chainguard/go AS builder
COPY . /app
RUN cd /app && go build -o app .

FROM cgr.dev/chainguard/glibc-dynamic
COPY --from=builder /app/app /usr/bin/
CMD ["/usr/bin/app"]
