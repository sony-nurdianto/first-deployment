FROM golang:1.24.4-bookworm

WORKDIR /ecommerce_web

COPY ecommerce/ecommerce_web/build/web .
COPY ecommerce/ecommerce_web/build/.env .
# RUN cp /ecommerce_web/web /usr/local/bin

CMD [ "./web" ]
