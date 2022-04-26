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

        stage('Build #1') {
            steps {
                echo 'Compiling and building'
                sh 'cd payment-mode && go build'
            }
        }

        stage('Test #1') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'cd payment-mode && go vet .'
                    echo 'Running test'
                    sh 'cd payment-mode && go test ./...'
                }
            }
        }

        stage('Create .env file #1') {
            steps {
                sh 'cd payment-mode && touch .env'
                sh 'cd payment-mode && echo "REGION=${REGION}" >> .env'
                sh 'cd payment-mode && echo "AWS_ACCESS_KEY_ID=${KEY_ID}" >> .env'
                sh 'cd payment-mode && echo "AWS_SECRET_ACCESS_KEY=${ACCESS_KEY}" >> .env'
                sh 'cd payment-mode && echo "SECRET=${JWT_SECRET_KEY}" >> .env'
            }
        }

        stage('Run #1') {
             steps {
                 echo 'Running the application'
                 sh 'cd payment-mode && chmod +x startup.sh && ./startup.sh'
             }
        }


        stage('Build #2') {
                    steps {
                        echo 'Compiling and building'
                        sh 'cd auth && go build'
                    }
                }

        stage('Test #2') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'cd auth && go vet .'
                    echo 'Running test'
                    sh 'cd auth && go test ./...'
                }
            }
        }

        stage('Create .env file #2') {
            steps {
                sh 'cd auth && touch .env'
                sh 'cd auth && echo "REGION=${REGION}" >> .env'
                sh 'cd auth && echo "AWS_ACCESS_KEY_ID=${KEY_ID}" >> .env'
                sh 'cd auth && echo "AWS_SECRET_ACCESS_KEY=${ACCESS_KEY}" >> .env'
                sh 'cd auth && echo "SECRET=${JWT_SECRET_KEY}" >> .env'
            }
        }

        stage('Run #2') {
             steps {
                 echo 'Running the application'
                 sh 'cd auth && chmod +x startup.sh && ./startup.sh'
             }
        }


        stage('Build #3') {
                    steps {
                        echo 'Compiling and building'
                        sh 'cd productsFrontStore && go build'
                    }
                }

        stage('Test #3') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'cd productsFrontStore && go vet .'
                    echo 'Running test'
                    sh 'cd productsFrontStore && go test ./...'
                }
            }
        }

        stage('Create .env file #3') {
            steps {
                sh 'cd productsFrontStore && touch .env'
                sh 'cd productsFrontStore && echo "REGION=${REGION}" >> .env'
                sh 'cd productsFrontStore && echo "AWS_ACCESS_KEY_ID=${KEY_ID}" >> .env'
                sh 'cd productsFrontStore && echo "AWS_SECRET_ACCESS_KEY=${ACCESS_KEY}" >> .env'
                sh 'cd productsFrontStore && echo "SECRET=${JWT_SECRET_KEY}" >> .env'
            }
        }

        stage('Run #3') {
             steps {
                 echo 'Running the application'
                 sh 'cd productsFrontStore && chmod +x startup.sh && ./startup.sh'
             }
        }

        stage('Build #4') {
                    steps {
                        echo 'Compiling and building'
                        sh 'cd productsAdmin && go build'
                    }
                }

        stage('Test #4') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'cd productsAdmin && go vet .'
                    echo 'Running test'
                    sh 'cd productsAdmin && go test ./...'
                }
            }
        }

        stage('Create .env file #4') {
            steps {
                sh 'cd productsAdmin && touch .env'
                sh 'cd productsAdmin && echo "REGION=${REGION}" >> .env'
                sh 'cd productsAdmin && echo "AWS_ACCESS_KEY_ID=${KEY_ID}" >> .env'
                sh 'cd productsAdmin && echo "AWS_SECRET_ACCESS_KEY=${ACCESS_KEY}" >> .env'
                sh 'cd productsAdmin && echo "SECRET=${JWT_SECRET_KEY}" >> .env'
            }
        }

        stage('Run #4') {
             steps {
                 echo 'Running the application'
                 sh 'cd productsAdmin && chmod +x startup.sh && ./startup.sh'
             }
        }

        stage('Build #5') {
                    steps {
                        echo 'Compiling and building'
                        sh 'cd order && go build'
                    }
                }

        stage('Test #5') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'cd order && go vet .'
                    echo 'Running test'
                    sh 'cd order && go test ./...'
                }
            }
        }

        stage('Create .env file #5') {
            steps {
                sh 'cd order && touch .env'
                sh 'cd order && echo "REGION=${REGION}" >> .env'
                sh 'cd order && echo "AWS_ACCESS_KEY_ID=${KEY_ID}" >> .env'
                sh 'cd order && echo "AWS_SECRET_ACCESS_KEY=${ACCESS_KEY}" >> .env'
                sh 'cd order && echo "SECRET=${JWT_SECRET_KEY}" >> .env'
            }
        }

        stage('Run #5') {
             steps {
                 echo 'Running the application'
                 sh 'cd order && chmod +x startup.sh && ./startup.sh'
             }
        }

        stage('Build #6') {
                    steps {
                        echo 'Compiling and building'
                        sh 'cd cart && go build'
                    }
                }

        stage('Test #6') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'cd cart && go vet .'
                    echo 'Running test'
                    sh 'cd cart && go test ./...'
                }
            }
        }

        stage('Create .env file #6') {
            steps {
                sh 'cd cart && touch .env'
                sh 'cd cart && echo "REGION=${REGION}" >> .env'
                sh 'cd cart && echo "AWS_ACCESS_KEY_ID=${KEY_ID}" >> .env'
                sh 'cd cart && echo "AWS_SECRET_ACCESS_KEY=${ACCESS_KEY}" >> .env'
                sh 'cd cart && echo "SECRET=${JWT_SECRET_KEY}" >> .env'
            }
        }

        stage('Run #6') {
             steps {
                 echo 'Running the application'
                 sh 'cd cart && chmod +x startup.sh && ./startup.sh'
             }
        }

        stage('Build #7') {
                    steps {
                        echo 'Compiling and building'
                        sh 'cd account-frontstore && go build'
                    }
                }

        stage('Test #7') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'cd account-frontstore && go vet .'
                    echo 'Running test'
                    sh 'cd account-frontstore && go test ./...'
                }
            }
        }

        stage('Create .env file #7') {
            steps {
                sh 'cd account-frontstore && touch .env'
                sh 'cd account-frontstore && echo "REGION=${REGION}" >> .env'
                sh 'cd account-frontstore && echo "AWS_ACCESS_KEY_ID=${KEY_ID}" >> .env'
                sh 'cd account-frontstore && echo "AWS_SECRET_ACCESS_KEY=${ACCESS_KEY}" >> .env'
                sh 'cd account-frontstore && echo "SECRET=${JWT_SECRET_KEY}" >> .env'
            }
        }

        stage('Run #7') {
             steps {
                 echo 'Running the application'
                 sh 'cd account-frontstore && chmod +x startup.sh && ./startup.sh'
             }
        }

        stage('Build #8') {
                    steps {
                        echo 'Compiling and building'
                        sh 'cd transaction && go build'
                    }
                }

        stage('Test #8') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'cd transaction && go vet .'
                    echo 'Running test'
                    sh 'cd transaction && go test ./...'
                }
            }
        }

        stage('Create .env file #8') {
            steps {
                sh 'cd transaction && touch .env'
                sh 'cd transaction && echo "REGION=${REGION}" >> .env'
                sh 'cd transaction && echo "AWS_ACCESS_KEY_ID=${KEY_ID}" >> .env'
                sh 'cd transaction && echo "AWS_SECRET_ACCESS_KEY=${ACCESS_KEY}" >> .env'
                sh 'cd transaction && echo "SECRET=${JWT_SECRET_KEY}" >> .env'
            }
        }

        stage('Run #8') {
             steps {
                 echo 'Running the application'
                 sh 'cd transaction && chmod +x startup.sh && ./startup.sh'
             }
        }

        stage('Build #9') {
                    steps {
                        echo 'Compiling and building'
                        sh 'cd gateway && go build'
                    }
                }

        stage('Test #9') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'cd gateway && go vet .'
                    echo 'Running test'
                    sh 'cd gateway && go test ./...'
                }
            }
        }

        stage('Create .env file #9') {
            steps {
                sh 'cd gateway && touch .env'
                sh 'cd gateway && echo "REGION=${REGION}" >> .env'
                sh 'cd gateway && echo "AWS_ACCESS_KEY_ID=${KEY_ID}" >> .env'
                sh 'cd gateway && echo "AWS_SECRET_ACCESS_KEY=${ACCESS_KEY}" >> .env'
                sh 'cd gateway && echo "SECRET=${JWT_SECRET_KEY}" >> .env'
            }
        }

        stage('Run #9') {
             steps {
                 echo 'Running the application'
                 sh 'cd gateway && chmod +x startup.sh && ./startup.sh'
             }
        }

        stage('Build #10') {
                    steps {
                        echo 'Compiling and building'
                        sh 'cd customer-admin && go build'
                    }
                }

        stage('Test #10') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'cd customer-admin && go vet .'
                    echo 'Running test'
                    sh 'cd customer-admin && go test ./...'
                }
            }
        }

        stage('Create .env file #10') {
            steps {
                sh 'cd customer-admin && touch .env'
                sh 'cd customer-admin && echo "REGION=${REGION}" >> .env'
                sh 'cd customer-admin && echo "AWS_ACCESS_KEY_ID=${KEY_ID}" >> .env'
                sh 'cd customer-admin && echo "AWS_SECRET_ACCESS_KEY=${ACCESS_KEY}" >> .env'
                sh 'cd customer-admin && echo "SECRET=${JWT_SECRET_KEY}" >> .env'
            }
        }

        stage('Run #10') {
             steps {
                 echo 'Running the application'
                 sh 'cd customer-admin && chmod +x startup.sh && ./startup.sh'
             }
        }

        stage('Build #11') {
                    steps {
                        echo 'Compiling and building'
                        sh 'cd shipping && go build'
                    }
                }

        stage('Test #11') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'cd shipping && go vet .'
                    echo 'Running test'
                    sh 'cd shipping && go test ./...'
                }
            }
        }

        stage('Create .env file #11') {
            steps {
                sh 'cd shipping && touch .env'
                sh 'cd shipping && echo "REGION=${REGION}" >> .env'
                sh 'cd shipping && echo "AWS_ACCESS_KEY_ID=${KEY_ID}" >> .env'
                sh 'cd shipping && echo "AWS_SECRET_ACCESS_KEY=${ACCESS_KEY}" >> .env'
                sh 'cd shipping && echo "SECRET=${JWT_SECRET_KEY}" >> .env'
            }
        }

        stage('Run #11') {
             steps {
                 echo 'Running the application'
                 sh 'cd shipping && chmod +x startup.sh && ./startup.sh'
             }
        }

        stage('Build #12') {
                    steps {
                        echo 'Compiling and building'
                        sh 'cd rewards && go build'
                    }
                }

        stage('Test #12') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'cd rewards && go vet .'
                    echo 'Running test'
                    sh 'cd rewards && go test ./...'
                }
            }
        }

        stage('Create .env file #12') {
            steps {
                sh 'cd rewards && touch .env'
                sh 'cd rewards && echo "REGION=${REGION}" >> .env'
                sh 'cd rewards && echo "AWS_ACCESS_KEY_ID=${KEY_ID}" >> .env'
                sh 'cd rewards && echo "AWS_SECRET_ACCESS_KEY=${ACCESS_KEY}" >> .env'
                sh 'cd rewards && echo "SECRET=${JWT_SECRET_KEY}" >> .env'
            }
        }

        stage('Run #12') {
             steps {
                 echo 'Running the application'
                 sh 'cd rewards && chmod +x startup.sh && ./startup.sh'
             }
        }

        stage('Build #13') {
                    steps {
                        echo 'Compiling and building'
                        sh 'cd checkout && go build'
                    }
                }

        stage('Test #13') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'cd checkout && go vet .'
                    echo 'Running test'
                    sh 'cd auth && go test ./...'
                }
            }
        }

        stage('Create .env file #13') {
            steps {
                sh 'cd checkout && touch .env'
                sh 'cd checkout && echo "REGION=${REGION}" >> .env'
                sh 'cd checkout && echo "AWS_ACCESS_KEY_ID=${KEY_ID}" >> .env'
                sh 'cd checkout && echo "AWS_SECRET_ACCESS_KEY=${ACCESS_KEY}" >> .env'
                sh 'cd checkout && echo "SECRET=${JWT_SECRET_KEY}" >> .env'
            }
        }

        stage('Run #13') {
             steps {
                 echo 'Running the application'
                 sh 'cd checkout && chmod +x startup.sh && ./startup.sh'
             }
        }

        stage('Build #14') {
                    steps {
                        echo 'Compiling and building'
                        sh 'cd categories && go build'
                    }
                }

        stage('Test #14') {
            steps {
                withEnv(["PATH+GO=${GOPATH}/bin"]){
                    echo 'Running vetting'
                    sh 'cd categories && go vet .'
                    echo 'Running test'
                    sh 'cd categories && go test ./...'
                }
            }
        }

        stage('Create .env file #14') {
            steps {
                sh 'cd categories && touch .env'
                sh 'cd categories && echo "REGION=${REGION}" >> .env'
                sh 'cd categories && echo "AWS_ACCESS_KEY_ID=${KEY_ID}" >> .env'
                sh 'cd categories && echo "AWS_SECRET_ACCESS_KEY=${ACCESS_KEY}" >> .env'
                sh 'cd categories && echo "SECRET=${JWT_SECRET_KEY}" >> .env'
            }
        }

        stage('Run #14') {
             steps {
                 echo 'Running the application'
                 sh 'cd categories && chmod +x startup.sh && ./startup.sh'
             }
        }
    }
}