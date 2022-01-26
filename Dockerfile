FROM scratch
COPY project2 ./
COPY config.env ./
EXPOSE 4003/tcp
EXPOSE 5003/tcp
ENTRYPOINT ["/project2"]