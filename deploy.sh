docker stop $1
docker rm $1
docker rmi $1
docker build -t $1:multistage -f Dockerfile.multistage .
docker run --publish 6969:6969 $1