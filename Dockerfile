FROM webstartdeps

# All folders related to the web app
COPY webapp/css/ /server/webapp/css
COPY webapp/img/ /server/webapp/img
COPY webapp/js/ /server/webapp/js
COPY webapp/views/ /server/webapp/views

# Setup environment
ENV GOPATH=/server/src/go
ENV SERVEROPT_DB_HOST "db:5432"

# Build the server
COPY . /server/src/go/src/github.com/anishmgoyal/webstart
WORKDIR /server/src/go/src/github.com/anishmgoyal/webstart
RUN go build && mv webstart /server/webstart

# Cleanup and final workdir
WORKDIR /server
RUN rm -rf /server/src

EXPOSE 2646

ENTRYPOINT [ "/server/webstart" ]