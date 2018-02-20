FROM alpine:3.7

ADD ./kibana-sidecar /kibana-sidecar

ENTRYPOINT ["/kibana-sidecar"]
CMD ["help"]
