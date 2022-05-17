FROM centos:7

EXPOSE 9090

ADD ./bin/run-platform-service /usr/bin/run-platform-service

CMD ["run-platform-service", "serve"]