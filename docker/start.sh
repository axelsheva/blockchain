cd 11.0.1.1
docker-compose build
cd ../11.0.2.1
docker-compose up -d --build
cd ../11.0.3.1
docker-compose up -d --build
cd ../11.0.1.1
docker-compose up --build
