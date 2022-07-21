FROM alpine
COPY ./ageCat ./ageCat
ENTRYPOINT ["./ageCat"]