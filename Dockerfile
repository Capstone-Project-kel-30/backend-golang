FROM python:3.8-slim
ARG port

USER root
COPY . /gym-app
WORKDIR /gym-app

ENV PORT=$port

RUN apt-get update && apt-get install -y --no-install-recommends apt-utils \
  && apt-get -y install curl \
  && apt-get install libgomp1

RUN chgrp -R 0 /gym-app \
  && chmod -R g=u /gym-app \
  && pip install pip --upgrade 
EXPOSE $PORT

CMD gunicorn app:server --bind 0.0.0.0:$PORT --preload