### Stage 1: Build server ###
FROM golang:alpine as go-builder
RUN apk update && apk upgrade && apk add --no-cache bash git openssh
# add credentials on build
#RUN git config --global url."https://github.com/".insteadOf "https://github.com/"
ARG token
RUN if [ "x$token" = "x" ] ; then echo "token not provided. trying build without it" ; else git config --global url."https://${token}:@github.com/".insteadOf "https://github.com/" ; fi

RUN mkdir /build 
ADD . /build/
WORKDIR /build 
ENV GIT_TERMINAL_PROMPT=1
RUN go build -o main .

### STAGE 2: Final container ###

FROM nginx:alpine

LABEL maintainer Igor Rendulic <igor@mail.io>

RUN apk --no-cache --update add ca-certificates

COPY nginx/nginx-operator.conf /etc/nginx/conf.d/nginx-operator.conf
RUN rm -rf /usr/share/nginx/html/* && rm /etc/nginx/conf.d/default.conf

COPY --from=go-builder /build/main /app/

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]