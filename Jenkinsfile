pipeline {
    // install golang 1.14 on Jenkins node
    agent any
    tools {
        go 'go1.22.4'
    }
    environment {
        DOCKERHUB_CREDENTIALS = credentials('telcsifarmer_dockerhub')
        IMAGE_NAME = "hajnalilepke/go_telcsifarmer"
        IMAGE_TAG = "latest"
        SSH_CREDENTIALS = 'telcsifarmer_ssh'
        REMOTE_HOST = 'malnacska@192.168.1.69'     
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
    }
    stages {
        stage('Checkout') {
            steps {
                // Checkout code from version control
                echo 'Checkout code from version control'
                checkout scm
            }
        }        
        stage("unit-test") {
            steps {
                echo 'UNIT TEST EXECUTION STARTED'
                // sh 'make unit-tests'
            }
        }
        stage("functional-test") {
            steps {
                echo 'FUNCTIONAL TEST EXECUTION STARTED'
                // sh 'make functional-tests'
            }
        }
        stage("build") {
            steps {
                echo 'BUILD EXECUTION STARTED'
                sh 'go version'
                sh 'go mod download'
                sh 'go build .'
                // sh 'docker build . go_telcsifarmer:multistage -f Dockerfile.multistage .'
            }
        }
        stage('Build Docker Image') {
            steps {
                script {
                    // Build the Docker image, assuming the Dockerfile is in the project root
                    docker.build("${IMAGE_NAME}:${IMAGE_TAG}")
                }
            }
        }
        stage('Push Docker Image') {
            steps {
                script {
                    docker.withRegistry('https://index.docker.io/v1/', DOCKERHUB_CREDENTIALS) {
                        docker.image("${IMAGE_NAME}:${IMAGE_TAG}").push()
                    }
                }
            }
        }        
        stage('Deploy') {
            steps {
                script {
                    // sshagent([SSH_CREDENTIALS]) {
                    //     // Trigger the script on the remote server to update the service
                        sh 'update_service.sh'
                    // }
                }
            }
        }
    }
}