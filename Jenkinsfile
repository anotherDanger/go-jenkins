pipeline{
    agent none
    environment{
            AUTHOR = "andhika danger"
            APP = credentials("app")
        }
    options{
        disableConcurrentBuilds()
        timeout(time: 1, unit: 'HOURS')
    }
   stages{
            stage('Information'){
                agent{
                    node{
                        label 'golang && almalinux'
                    }
                }

                steps{
                    echo("Author: ${AUTHOR}")
                    echo("App: ${APP_USR}")
                    sh('echo "APP PSW: $APP_PSW" > "rahasia.txt"')
                    echo("Job Name: ${env.JOB_NAME}")
                    echo("Node Labels: ${env.NODE_LABELS}")
                    echo("Branch Name: ${env.BRANCH_NAME}")
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