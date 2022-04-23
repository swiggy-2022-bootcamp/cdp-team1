#!groovy
pipeline {
    agent any
    tools {
        go 'Go'
    }
    environment {
        GO111MODULE = 'on'
        CGO_ENABLED = 0
        GOPATH = "${JENKINS_HOME}/jobs/${JOB_NAME}/builds/${BUILD_ID}"
        APP_NAME = "payment-mode"
        REGION = "${REGION}"
        ACCESS_KEY = "${AWS_SECRET_ACCESS_KEY}"
        KEY_ID = "${AWS_ACCESS_KEY_ID}"
        JWT_SECRET_KEY = "${JWT_SECRET_KEY}"
    }
    stages {
        stage('Pre Test') {
            steps {
                echo 'Installing dependencies'
                sh 'go version'
            }
        }

        stage('Build') {
            steps {
                echo 'Compiling and building'
                sh 'cd payment-mode && go build'
            }
        }

        stage('Test') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'cd payment-mode && go vet .'
                    echo 'Running test'
                    sh 'cd payment-mode && go test ./...'
                }
            }
        }

        stage('Create .env file') {
            steps {
                sh 'cd payment-mode && touch .env'
                sh 'cd payment-mode && echo "REGION=${REGION}" >> .env'
                sh 'cd payment-mode && echo "AWS_ACCESS_KEY_ID=${KEY_ID}" >> .env'
                sh 'cd payment-mode && echo "AWS_SECRET_ACCESS_KEY=${ACCESS_KEY}" >> .env'
            }
        }

        stage('Run') {
             steps {
                 echo 'Running the application'
                 sh 'cd payment-mode && chmod +x startup.sh && ./startup.sh'
             }
        }
    }
}