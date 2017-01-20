FROM scratch
LABEL name="api"
ENV GIN_MODE release
ADD dist/server /
CMD ["/server"]
EXPOSE 3000
