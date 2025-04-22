pipeline{
    agent any
    stages{
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

        stage('Unit test'){
            steps{
                sh 'docker run --rm my-app go test ./repository ./service ./controller'
            }
        }

        stage('Curl'){
            steps{
                sh 'curl http://202.10.36.121:8081/v1/book'
            }
        }

        stage('Clear'){
            steps{
                sh 'docker compose down'
            }
        }
    }
}