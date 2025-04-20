pipeline{
   agent{
    node{
        labels 'golang && almalinux'
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
}