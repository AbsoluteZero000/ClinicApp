docker build -t backend ./Backend/awesomeProject

docker build -t database ./Database

docker build -t frontend ./Frontend/app

docker network create mynetwork

docker run --name database --network=mynetwork -e MYSQL_ROOT_PASSWORD=12345678 -d database:latest
sleep 5
docker run --name backend-container -e DB_DRIVER=mysql -e DB_USER=root -e DB_PASS=12345678 -e DB_NAME=ClinicApp -p 8080:8080 --network=mynetwork -d backend:latest

docker run --name frontend -e apiKey=http://localhost:8080  --network=mynetwork -p 4200:4200 -d frontend:latest
