FROM alpine:3.5

COPY bin/binary /usr/local/bin/app

# RUN apk --update add docker

# https://serverfault.com/questions/772227/chmod-not-working-correctly-in-docker
RUN chmod +x /usr/local/bin/app

ENTRYPOINT []
CMD [ "app", "server" ]