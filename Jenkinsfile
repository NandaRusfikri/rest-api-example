pipeline {
  agent any

  environment {
    PATH = "/usr/local/go/bin:${env.PATH}"
  }

  stages {
    stage('Check Requirements') {
      steps {
        echo 'Pulling... ' + env.GIT_BRANCH

        script {
          // Retrieve the last commit message using git command
          env.GIT_COMMIT_MSG = sh(script: 'git log -1 --pretty=%B ${GIT_COMMIT}', returnStdout: true).trim()
        }

        echo "${GIT_COMMIT_MSG}"
        sh 'whoami'
        echo 'Installing dependencies'
        sh 'which go'
        sh 'go version'
      }
    }

    stage('Run Test') {
      steps {
        echo 'Belum Ada Test silahkan lanjut'
      }
    }

    stage('Build Binary') {
      steps {
        echo 'Compiling and building Binary'

        sh 'CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags="-s -w" -o rest-api main.go'
        sh 'chmod +x rest-api'
      }
    }

    stage('Deploy Server Development') {
      steps {
        sh 'scp rest-api root@167.71.206.43:/root/rest-api'
        sh 'ssh root@167.71.206.43'
        sh 'pwd'
        sh 'cd /root/rest-api'
        sh 'fuser -n tcp -k 54321'

        sh 'nohup ./rest-api > api.log &'
      }
    }




  }
}
