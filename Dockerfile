FROM cydev/go
RUN go get "github.com/ernado/gomp/gompd"
EXPOSE 80
VOLUME /data
WORKDIR /data
ENTRYPOINT ["gompd"]
CMD["-port", "80",]
