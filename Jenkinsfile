pipeline{
    agent none
    environment{
            AUTHOR = "andhika danger"
            APP = credentials("app")
        }
    parameters{
        text(name: 'NAME', defaultValue: '', description: 'Describe who you are')
        booleanParam(name: 'DEPLOY', defaultValue: false, description: 'Need to deploy?')
        choice(name: 'SOCIAL', choices: ['Facebook', 'Instagram', 'X'], description: 'Social Media')
    }
    triggers{
        cron('* * * * *')
    }
   stages{
            stage('Intro'){
                agent{
                    node{
                        label 'golang && almalinux'
                    }
                }

                steps{
                    echo("Hello ${params.NAME}")
                    echo("Time to deploy? ${params.DEPLOY}")
                    echo("Your social media is ${params.SOCIAL}")
                }
            }
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