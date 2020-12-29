FROM golang:buster as BUILD
WORKDIR /build
COPY . /build
RUN go build .

FROM python:3.9.1-slim-buster
ENV PYTHONPATH=/app
WORKDIR /app
COPY text_stemmer/ /app
COPY start.sh /app
RUN python3 -m pip install -r requirements.txt
COPY --from=BUILD /build/feiqineus .
EXPOSE 8080
CMD ["sh", "start.sh"]
