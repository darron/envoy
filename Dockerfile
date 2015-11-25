# https://hub.docker.com/r/darron/envoy/
FROM octohost/nginx:1.8

RUN curl -sfL -o envoy-0.3-linux-amd64.gz https://github.com/darron/envoy/releases/download/v0.3/envoy-0.3-linux-amd64.gz \
  && gunzip envoy-0.3-linux-amd64.gz \
  && mv envoy-0.3-linux-amd64 /usr/local/bin/envoy \
  && chmod 755 /usr/local/bin/envoy
