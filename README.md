GO DEPLOY
docker build -t go_telcsifarmer:multistage -f Dockerfile.multistage .
docker run --publish 6969:6969 go_telcsifarmer:multistage

JENKINS
docker run -p 8080:8080 -p 50000:50000 -v jenkins_home:/var/jenkins_home jenkins/jenkins:latest

NEW WAY
docker compose up
