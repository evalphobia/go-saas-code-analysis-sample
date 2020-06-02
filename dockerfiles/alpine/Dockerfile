FROM python:3.8.3-alpine

RUN apk add --update curl haproxy sqlite-dev gcc g++ \
    && rm -rf /var/cache/apk/*

RUN ln -s /usr/include/locale.h /usr/include/xlocale.h

RUN pip install --upgrade pip

RUN pip install boto bottle numpy cython pandas oauth2client

EXPOSE 11111 11111

ENTRYPOINT ["python", "-c", "\"print('hello')\""]
