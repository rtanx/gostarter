# syntax=docker/dockerfile:1

#dependencies
FROM golang:1.18-alpine

# RUN apk add --no-cache libstdc++ libx11 libxrender libxext libssl1.1 ca-certificates \
#     && apk add --no-cache --virtual .build-deps \
#     msttcorefonts-installer \
#     # Install microsoft fonts
#     && update-ms-fonts \
#     && fc-cache -f \
#     # Clean up when done
#     && rm -rf /tmp/* \
#     && apk del .build-deps


ARG ENV
ENV ENV ${ENV}
WORKDIR /app

COPY . ./

RUN go mod download
# RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN if [[ ${ENV} == "DEV" ]] || [[ ${ENV} == "dev" ]]; then \
      apk update; \
      apk add --no-cache make openssh-client curl gcc libc-dev; \
      curl -fLo install.sh https://raw.githubusercontent.com/cosmtrek/air/master/install.sh \
      && chmod +x install.sh && sh install.sh && cp ./bin/air /bin/air; \
    else \
      go build -o /app; \
    fi

EXPOSE 8000

CMD if [[ ${ENV} == "DEV" ]] || [[ ${ENV} == "dev" ]]; then air; else /app start; fi
