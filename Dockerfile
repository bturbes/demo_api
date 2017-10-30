FROM debian:stable-slim

COPY bin/demo_api.linux /usr/local/bin/demo_api

CMD [ "/usr/local/bin/demo_api" ]
