# Subscriber for NATS-streaming
___
## Contents
* [General info](#general-info)
* [Technologies](#technologies)
* [Setup](#setup)
* [Example](#example)

## General info
Simple service for saving orders from NATS-streaming to Postgres, checking format of incoming message and stream it to 
browser. Also messages savings to cache and service provide to get row by id from cache.

In directory "Produser" realised simple NATS sender script. It parce model.json file and send bytes to subscriber ones 
on 5 seconds.

## Technologies
Project is created with:
* GO
* Postgres
* Docker
* GORM
* SSE
* NATS-streaming
* Javascript

## Setup

API uses ports:

* 8080
* 5432

Up NATS streaming server locally and docker-compose up. 
Streaming starts on localhost:8080/.

Get order by id from cache you can on localhost:8080/orders/id



