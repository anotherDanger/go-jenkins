pipeline{
   agent{
    node{
        label "golang && almalinux"
    }
   }
   stages{
            stage('Build'){
                steps{
                    sh 'go build -o main'
                }
            }
            stage('Test'){
                steps{
                    sh 'go test ./repository ./service ./controller -cover'
                }
            }
        }
    post{
        always{
            echo("Starting........")
        }

        success{
            echo("Completed, and success boss!")
        }

        failure{
            echo("Build failed")
        }

        cleanup{
            echo("End.")
        }
    }
}