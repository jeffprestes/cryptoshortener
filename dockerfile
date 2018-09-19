#env GOOS=linux GOARCH=amd64 go build
#docker build -t mercurius:cryptoshortner .
#docker run -p 8080:8080 -d mercurius:cryptoshortner

FROM scratch

ADD cryptoshortner /
ADD conf/ /conf
ADD public/ /public
ADD locale/ /locale

CMD [ "/cryptoshortner" ]