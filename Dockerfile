FROM centos
MAINTAINER mingz2013
COPY magic-cube-go /sbin/magic-cube-go
RUN chmod 755 /sbin/magic-cube-go
ENTRYPOINT ["/sbin/magic-cube-go"]
