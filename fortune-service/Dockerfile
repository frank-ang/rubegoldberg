FROM alpine:3.16

# Update
RUN apk update && apk upgrade

# Install dependencies if any

## XRay?
#RUN apk add libc6-compat
#RUN rm -rf /tmp/xray.rpm
#RUN curl https://s3.ap-southeast-1.amazonaws.com/aws-xray-assets.ap-southeast-1/xray-daemon/aws-xray-daemon-linux-3.x.zip -o /tmp/xray.zip
# ...

# Install app
RUN rm -rf /opt/fortune/*
RUN mkdir -p /opt/fortune
ADD fortune-linux /opt/fortune
ENV FOO "bar"

EXPOSE 80

CMD ["/opt/fortune/fortune-linux", "-mysql", "-redis", "-es"]
