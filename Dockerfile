## Build telegram bot
FROM golang:buster as BUILD
WORKDIR /build
COPY . /build
RUN go build .

## Run time for python app
FROM python:3-slim
ENV PYTHONPATH=/app
WORKDIR /app

## Copy telegram bot and stemmer
COPY --from=BUILD /build/feiqineus .
COPY text_stemmer/ /app
COPY start.sh /app

# Install requirements
RUN pip install --upgrade pip
RUN pip install pipenv
RUN pipenv lock --keep-outdated --requirements > requirements.txt
RUN python3 -m pip install -r /app/requirements.txt --no-cache-dir
EXPOSE 8080
CMD ["sh", "start.sh"]
