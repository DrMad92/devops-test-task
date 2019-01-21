FROM alpine:latest

ENV SERVER "localhost"
ENV PORT "5432"
ENV HOST "0.0.0.0:8000"
ENV USER "testdbmaster"
ENV PASSWORD "testdbmasterpass"
ENV DBNAME "testdb1"

ADD main /
ADD template.tmpl /

ENTRYPOINT /main -dbserver ${SERVER} -host ${HOST} -port ${PORT} -user ${USER} -password ${PASSWORD} -dbname ${DBNAME}