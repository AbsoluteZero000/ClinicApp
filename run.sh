docker build -t backend ./Backend/awesomeProject

docker build -t database ./Database

docker build -t frontend ./Frontend/app

docker network create mynetwork

docker run --name mysql --network=mynetwork -e MYSQL_ROOT_PASSWORD=12345678 -d database:latest
sleep 5

docker run --name backend -e DB_DRIVER=mysql -e DB_USER=root -e DB_PASS=12345678 -e DB_NAME=ClinicApp -p 8080:8080 --network=mynetwork -d backend:latest

docker run --name frontend  --network=mynetwork -p 4200:4200 -d frontend:latest
