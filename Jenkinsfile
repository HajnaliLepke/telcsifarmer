pipeline {
    // install golang 1.14 on Jenkins node
    agent any
    tools {
        go 'go1.22.4'
    }
    environment {
        GO114MODULE = 'on'
        CGO_ENABLED = 0 
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
    }
    stages {
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
                sh 'go get ./...'
                sh 'docker stop go_telcsifarmer'
                sh 'docker rm go_telcsifarmer'
                sh 'docker rmi go_telcsifarmer:multistage'
                sh 'docker build . go_telcsifarmer:multistage -f Dockerfile.multistage .'
            }
        }
        stage('deliver') {
            agent any
            steps {
                sh 'docker run --publish 6969:6969 go_telcsifarmer:multistage'
            }
        }
    }
}