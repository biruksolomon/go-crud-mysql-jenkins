pipeline {
    agent any

    environment {
        GO_VERSION = "1.21"
        IMAGE_NAME = "go-crud-api"
    }

    stages {

        stage('Checkout Code') {
            steps {
                echo 'Pulling code from GitHub...'
                checkout scm
            }
        }

        stage('Install Dependencies') {
            steps {
                echo 'Installing Go dependencies...'
                sh 'go mod tidy'
            }
        }

        stage('Run Tests') {
            steps {
                echo 'Running tests...'
                sh 'go test ./...'
            }
        }

        stage('Build Go Application') {
            steps {
                echo 'Building Go binary...'
                sh 'go build -o app'
            }
        }

        stage('Build Docker Image') {
            steps {
                echo 'Building Docker image...'
                sh 'docker build -t $IMAGE_NAME .'
            }
        }

    }

    post {
        success {
            echo '✅ Pipeline completed successfully!'
        }
        failure {
            echo '❌ Pipeline failed!'
        }
    }
}
