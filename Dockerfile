FROM scratch
MAINTAINER jspc <james.condron@ft.com>

COPY ./awesomesul /awesomesul

EXPOSE 8000
ENTRYPOINT ["/awesomesul"]
