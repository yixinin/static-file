FROM centos:7

COPY static-file /app/static-file

WORKDIR /app

CMD ["/app/static-file", "-c", "/app/config/config.yaml"]