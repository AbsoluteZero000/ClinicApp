docker run --name mysql-container --network=mynetwork -e MYSQL_ROOT_PASSWORD=12345678 -d your_mysql_image:latest

docker run --name backend-container --network=mynetwork -p 8080:8080 -d backend:latest

docker run --name frontend --network=mynetwork -p 4200:4200 -d frontend:latest
