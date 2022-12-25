FROM python:3.8-slim-buster
COPY requirements.txt requirements.txt
RUN pip install -r requirements.txt
WORKDIR /backend
COPY . .
CMD python nlp.py