pipeline {
    agent any

    environment {
        DOCKERHUB_CREDENTIALS = credentials('dockerhub-creds')
        IMAGE_NAME     = "rashilmanandhar/go-hello-world"
        CONTAINER_NAME = "go-hello-world-test"
    }

    stages {

        stage('Clean workspace') {
            steps {
                cleanWs()
            }
        }

        stage('Checkout') {
            steps {
                git branch: 'main',
                    url: 'https://github.com/rashilsayami/jekins.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                sh "docker build -t ${IMAGE_NAME}:${BUILD_NUMBER} ."
            }
        }

        stage('List Docker Images') {
            steps {
                sh 'docker images'
            }
        }

        stage('Remove Old Test Container (if any)') {
            steps {
                sh "docker rm -f ${CONTAINER_NAME} || true"
            }
        }

        stage('Run Container (test)') {
            steps {
                sh "docker run -d -p 9090:8080 --name ${CONTAINER_NAME} ${IMAGE_NAME}:${BUILD_NUMBER}"
            }
        }

        stage('Wait 60s') {
            steps {
                sleep(time: 60, unit: 'SECONDS')
            }
        }

        stage('Stop & Remove Test Container') {
            steps {
                sh "docker stop ${CONTAINER_NAME}"
                sh "docker rm ${CONTAINER_NAME}"
            }
        }

        stage('Push to DockerHub') {
            steps {
                sh "echo \$DOCKERHUB_CREDENTIALS_PSW | docker login -u \$DOCKERHUB_CREDENTIALS_USR --password-stdin"
                sh "docker push ${IMAGE_NAME}:${BUILD_NUMBER}"
            }
        }

    }

    post {
        always {
            echo "Build Number: ${BUILD_NUMBER}"
            sh 'docker logout'
        }
    }
}
