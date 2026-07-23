pipeline {
    agent any

    triggers {
        pollSCM('H/3 * * * *')   // checks GitHub for changes every ~3 min
    }

    environment {
        DOCKERHUB_CREDENTIALS = credentials('dockerhub-creds')
        IMAGE_NAME     = "rashilmanandhar/go-hello-world"
        CONTAINER_NAME = "go-hello-world-test"
        DEPLOY_DIR     = "/opt/go-hello-world"   // where docker-compose.yml lives on the server
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
                // push the unique build-numbered tag
                sh "docker push ${IMAGE_NAME}:${BUILD_NUMBER}"
                // also tag/push as 'latest' so the server always pulls the newest one
                sh "docker tag ${IMAGE_NAME}:${BUILD_NUMBER} ${IMAGE_NAME}:latest"
                sh "docker push ${IMAGE_NAME}:latest"
            }
        }

        stage('Deploy to Server') {
            steps {
                sh """
                    cd ${DEPLOY_DIR}
                    docker compose pull
                    docker compose up -d --force-recreate
                """
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
