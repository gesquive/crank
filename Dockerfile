FROM gesquive/go-builder:latest AS builder

ENV APP=crank

# This requires that `make release-snapshot` be called first
COPY dist/ /dist/
RUN copy-release
RUN chmod +x /app/crank

# =============================================================================
FROM gesquive/docker-base:latest
LABEL maintainer="Gus Esquivel <gesquive@gmail.com>"

COPY --from=builder /app/${APP} /app/

ENTRYPOINT ["/app/crank"]
