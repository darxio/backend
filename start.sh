cd /home/artem/darxio/backend
git checkout $1
git pull 
docker-compose up --build -d > /dev/null


