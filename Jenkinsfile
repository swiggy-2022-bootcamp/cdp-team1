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
        APP_NAME = "shipping"
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
                    sh 'cd shipping && go vet .'
                    echo 'Running test'
                    sh 'cd shipping && go test ./...'
                }
            }
        }

        stage('Run') {
             steps {
                 echo 'Running the application'
                 sh 'cd shipping && ./shipping'
             }
        }
    }
}