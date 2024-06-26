#!/bin/bash

# # Change to the directory containing the docker-compose.yml file
# cd /path/to/your/docker-compose-file

# Pull the latest image
docker-compose pull

# Recreate and restart the service
docker-compose up -d
