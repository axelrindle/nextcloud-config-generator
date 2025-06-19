ARG NEXTCLOUD_VERSION=31.0.5

FROM nextcloud:${NEXTCLOUD_VERSION}

COPY custom-entrypoint.sh /

RUN chmod +x /custom-entrypoint.sh

ENV NEXTCLOUD_UPDATE=1
ENTRYPOINT []
CMD [ "/custom-entrypoint.sh" ]
