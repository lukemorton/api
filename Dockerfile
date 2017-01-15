FROM scratch
LABEL name="api"
ADD dist/server /
CMD ["/server"]
EXPOSE 3000
