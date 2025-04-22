pipeline{
    agent any
    stages{
        stage('compose down'){
            steps{
                sh 'docker compose down'
            }
        }
        stage('Build dockerfile'){
            steps{
                sh 'docker build -t my-app .'
            }
        }

        stage('Docker compose'){
            steps{
                sh 'docker compose up -d'
                script{
                    sleep(time: 5, unit: 'SECONDS')
                }
                sh 'docker compose ps'
            }
        }
}
}
