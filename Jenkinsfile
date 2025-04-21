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
            stage('Deploy'){
                input{
                    message: "Can we deploy?"
                    ok: "Yessir!"
                    submitter: "anotherDanger"
                    parameters{
                        string(name: 'QA', defaultValue: 'Anonymous', description: 'Who r u?')
                        choice(name: 'TARGET_ENV', choices: ['Development', 'Staging', 'Production'], description: 'Choose environment to deploy')
                    }
                    }
                 agent{
                    node{
                        label 'golang && almalinux'
                    }
                }
                steps{
                    echo("Deployed to ${TARGET_ENV}, approved by ${QA}")
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