GO DEPLOY
docker build -t go_telcsifarmer:multistage -f Dockerfile.multistage .
docker run --publish 6969:6969 go_telcsifarmer:multistage
