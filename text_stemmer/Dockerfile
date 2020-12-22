FROM python:3.9.1-buster
COPY requirements.txt .
RUN python3 -m pip install -r requirements.txt
ENV PYTHONPATH=/app/
COPY *.py ./app/
COPY nltk_data ./app/nltk_data
WORKDIR ./app/
EXPOSE 8081
CMD ["python3", "main.py"]