pipeline{
    agent none
    environment{
            AUTHOR = "andhika danger"
            APP = credentials("app")
        }
    // triggers{
    //     // cron('* * * * *')
    //     pollSCM('* * * * *')
    // }
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
             stage('Example') {
            input {
                message "Should we continue?"
                ok "Yes, we should."
                submitter "anotherDanger"
                parameters {
                    string(name: 'PERSON', defaultValue: 'Mr Jenkins', description: 'Who should I say hello to?')
                }
            }
            steps {
                echo "Hello, ${PERSON}, nice to meet you."
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