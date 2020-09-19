FROM plugins/base:multiarch

LABEL maintainer="Robert Kaussow <mail@thegeeklab.de>" \
    org.label-schema.name="Drone GitHub Comment" \
    org.label-schema.vendor="Robert Kaussow" \
    org.label-schema.schema-version="1.0"

ADD release/drone-github-release /bin/
ENTRYPOINT [ "/bin/drone-github-release" ]
