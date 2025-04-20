pipeline{
    agent none
   stages{
            stage('Information'){
                agent{
                    node{
                        label 'golang && almalinux'
                    }
                }

                steps{
                    echo("Job Name: ${JOB_NAME}")
                    echo("Node Labels: ${NODE_LABELS}")
                    echo("Branch Name: ${BRANCH_NAME}")
                }
            }
            stage('Build'){
                agent{
                    node{
                        label 'golang && almalinux'
                    }
                }
                steps{
                    sh 'go build -o main'
                }
            }
            stage('Test'){
                agent {
                    node{
                        label 'golang && almalinux'
                    }
                }
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