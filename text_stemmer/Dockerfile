FROM python:3-slim

# Install requirements
RUN pip install --upgrade pip
RUN pip install pipenv
COPY Pipfile* /tmp
RUN cd /tmp && pipenv lock --keep-outdated --requirements > requirements.txt
RUN python3 -m pip install -r /tmp/requirements.txt --no-cache-dir

## Copy codefiles
RUN mkdir /app
COPY *.py /app
COPY nltk_data /app/nltk_data

## Run configs
WORKDIR /app/
ENV PYTHONPATH=/app/
EXPOSE 8000
CMD ["uvicorn", "main:app", "--host", "0.0.0.0"]