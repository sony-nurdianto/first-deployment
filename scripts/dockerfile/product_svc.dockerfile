FROM golang:1.24.4-bookworm

WORKDIR /product_svc

COPY ecommerce/product_svc/build/service service
COPY ecommerce/product_svc/build/.env .env
# RUN cp /product_svc/service /usr/local/bin

CMD [ "./service" ]
