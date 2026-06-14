FROM debian:stable-slim

RUN apt update
RUN apt upgrade -y

RUN apt install -y build-essential zlib1g-dev libncurses5-dev libgdbm-dev libnss3-dev libssl-dev libreadline-dev libffi-dev libsqlite3-dev wget libbz2-dev

RUN wget https://www.python.org/ftp/python/3.10.8/Python-3.10.8.tgz
RUN tar -xf Python-3.10.*.tgz

RUN cd Python-3.10.8 && ./configure --enable-optimizations && make && make altinstall

COPY main.py main.py

COPY books/ books/

CMD ["python3.10", "main.py"]
