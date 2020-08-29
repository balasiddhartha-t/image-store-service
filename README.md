# Image Management System

An Image Management System built using golang.

It uses the File System Persistance for Create/delete Images into Albums.

## Pre requistes:
    - docker
    - docker-compose

## Application:

For starting the application you can run 

#### docker-compose build

This will download the necessary files required for the application to be started

#### docker-compose up -d

Above line will make sure that all the containers are created in detached mode and run in the background.

Note: In the docker-compose.yaml file, Under the go-image-store-service please mount volume to the path where you want the Albums to be stored.

Once the containers are up, you can read the information related to the end points at

###### http://localhost:8081/docs

