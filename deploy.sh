# docker stop go_telcsifarmer
# docker rm go_telcsifarmer
# docker rmi go_telcsifarmer:multistage
docker build -t go_telcsifarmer:multistage -f Dockerfile.multistage .
docker run --publish 6969:6969 --name go_telcsifarmer go_telcsifarmer:multistage