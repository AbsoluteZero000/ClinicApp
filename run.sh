docker build -t backend /Backend/awesomeProject

docker build -t database /Database

docker build -t frontend /Frontend/app


docker run --name database --network=mynetwork -e MYSQL_ROOT_PASSWORD=12345678 -d database:latest
sleep 5
docker run --name backend --network=mynetwork -p 8080:8080 -d backend:latest

docker run --name frontend --network=mynetwork -p 4200:4200 -d frontend:latest
