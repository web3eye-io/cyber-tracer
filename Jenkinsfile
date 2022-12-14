pipeline {
  agent any
  environment {
    GOPROXY = 'https://goproxy.cn,direct'
  }
  tools {
    go 'go'
  }
  stages {
    stage('Clone') {
      steps {
        git(url: scm.userRemoteConfigs[0].url,credentialsId: 'jiangjie-git-ssh-private-key', branch: '$BRANCH_NAME', changelog: true, poll: true)
      }
    }

    stage('Prepare') {
      steps {
        sh 'make deps'
      }
    }

    stage('Linting') {
      when {
        expression { BUILD_TARGET == 'true' }
      }
      steps {
        sh 'make verify'
      }
    }

    stage('Compile') {
      when {
        expression { BUILD_TARGET == 'true' }
      }
      steps {
        sh (returnStdout: false, script: '''
          make verify-build
        '''.stripIndent())
        sh 'DOCKER_REGISTRY=$DOCKER_REGISTRY make build-docker'
      }
    }

    // TODO: support switch k8s cluster
    // stage('Switch to current cluster') {
    //   when {
    //     anyOf {
    //       expression { BUILD_TARGET == 'true' }
    //       expression { DEPLOY_TARGET == 'true' }
    //     }
    //   }
    //   steps {
    //     sh 'cd /etc/kubeasz; ./ezctl checkout $TARGET_ENV'
    //   }
    // }

    // TODO: support UT
    // stage('Unit Tests') {
    //   when {
    //     expression { BUILD_TARGET == 'true' }
    //   }
    //   steps {
    //     sh (returnStdout: false, script: '''
    //       devboxpod=`kubectl get pods -A | grep development-box | head -n1 | awk '{print $2}'`
    //       servicename="nft-meta"

    //       kubectl exec --namespace kube-system $devboxpod -- make -C /tmp/$servicename after-test || true
    //       kubectl exec --namespace kube-system $devboxpod -- rm -rf /tmp/$servicename || true
    //       kubectl cp ./ kube-system/$devboxpod:/tmp/$servicename

    //       kubectl exec --namespace kube-system $devboxpod -- make -C /tmp/$servicename deps before-test test after-test
    //       kubectl exec --namespace kube-system $devboxpod -- rm -rf /tmp/$servicename

    //       swaggeruipod=`kubectl get pods -A | grep swagger | awk '{print $2}'`
    //       kubectl cp message/npool/*.swagger.json kube-system/$swaggeruipod:/usr/share/nginx/html || true
    //     '''.stripIndent())
    //   }
    // }
    stage('Tag') {
      when {
        anyOf{
          expression { TAG_MAJOR == 'true' }
          expression { TAG_MINOR == 'true' }
          expression { TAG_PATCH == 'true' }
        }
        anyOf{
          expression { TAG_FOR == 'test' }
          expression { TAG_FOR == 'prod' }
        }
      }
      steps {
        sh(returnStdout: true, script: '''
          set +e
          revlist=`git rev-list --tags --max-count=1`
          rc=$?
          set -e

          major=0
          minor=0
          patch=-1
          
          if [ 0 -eq $rc ]; then
            tag=`git describe --tags $revlist`

            major=`echo $tag | awk -F '.' '{ print $1 }'`
            minor=`echo $tag | awk -F '.' '{ print $2 }'`
            patch=`echo $tag | awk -F '.' '{ print $3 }'`

            if [ "$TAG_MAJOR" == 'true' ]; then
              major=$(( $major + 1 ))
              minor=0
              patch=-1
            elif [ "$TAG_MINOR" == 'true' ]; then
              minor=$(( $minor + 1 ))
              patch=-1
            fi    
          fi

          case $TAG_FOR in
            test)
              patch=$(( $patch + $patch % 2 + 1 ))
              ;;
            prod)
              patch=$(( $patch + ( $patch +  1 ) % 2 + 1 ))
              git reset --hard
              git checkout $tag
              ;;
          esac

          tag=$major.$minor.$patch
          
          git tag -a $tag -m "Bump version to $tag"
        '''.stripIndent())

        withCredentials([gitUsernamePassword(credentialsId: 'jiangjie-git-username-passwd', gitToolName: 'git-tool')]) {
          sh 'git push --tag'
        }
      }
    }
    
    stage('Generate docker image for development') {
      when {
        expression { RELEASE_TARGET == 'true' }
        expression { TAG_FOR == 'dev' }
      }
      steps {
        sh 'DEVELOPMENT=development DOCKER_REGISTRY=$DOCKER_REGISTRY make build-docker'
        sh 'TAG=latest DOCKER_REGISTRY=$DOCKER_REGISTRY make release-docker'
      }
    }

    stage('Generate docker image for feature test') {
      when {
        expression { RELEASE_TARGET == 'true' }
        expression { TAG_PATCH == 'true' }
        expression { TAG_FOR == 'test' }
      }
      steps {
        sh(returnStdout: false, script: '''
          revlist=`git rev-list --tags --max-count=1`
          tag=`git describe --tags $revlist`

          set +e
          docker images  | grep $tag
          rc=$?
          set -e
          if [ ! 0 -eq $rc ]; then
            DEVELOPMENT=test DOCKER_REGISTRY=$DOCKER_REGISTRY make build-docker
          fi

          set +e
          docker images  | grep $tag
          rc=$?
          set -e
          if [ 0 -eq $rc ]; then
            TAG=$tag DOCKER_REGISTRY=$DOCKER_REGISTRY make release-docker
          fi
        '''.stripIndent())
        sh ''
      }
    }

    stage('Generate docker image for testing or production') {
      when {
        anyOf {
          expression { RELEASE_TARGET == 'true' }
          expression { TAG_PATCH == 'true' }
          expression { TAG_FOR == 'prod' }
        }
      }
      steps {
          sh(returnStdout: false, script: '''
          revlist=`git rev-list --tags --max-count=1`
          tag=`git describe --tags $revlist`

          major=`echo $tag | awk -F '.' '{ print $1 }'`
          minor=`echo $tag | awk -F '.' '{ print $2 }'`
          patch=`echo $tag | awk -F '.' '{ print $3 }'`

          patch=$(( $patch - $patch % 2 ))
          tag=$major.$minor.$patch

          set +e
          docker images | grep $tag
          rc=$?
          set -e
          if [ ! 0 -eq $rc ]; then
            DEVELOPMENT=prod DOCKER_REGISTRY=$DOCKER_REGISTRY make build-docker
          fi

          set +e
          docker images | grep $tag
          rc=$?
          set -e
          if [ 0 -eq $rc ]; then
            TAG=$tag DOCKER_REGISTRY=$DOCKER_REGISTRY make release-docker
          fi
        '''.stripIndent())
      }
    }

    stage('Deploy for development') {
      when {
        expression { DEPLOY_TARGET == 'true' }
        expression { TARGET_ENV == "dev" }
      }
      steps {
        sh 'TAG=latest  make deploy-to-k8s-cluster'
      }
    }

    stage('Deploy for testing') {
      when {
        expression { DEPLOY_TARGET == 'true' }
        expression { TARGET_ENV == "test" }
      }
      steps {
        sh(returnStdout: true, script: '''
          revlist=`git rev-list --tags --max-count=1`
          tag=`git describe --tags $revlist`

          git reset --hard
          git checkout $tag
          TAG=$tag make deploy-to-k8s-cluster
        '''.stripIndent())
      }
    }

    stage('Deploy for production') {
      when {
        expression { DEPLOY_TARGET == 'true' }
        expression { TARGET_ENV == "prod" }
      }
      steps {
        sh(returnStdout: true, script: '''
          revlist=`git rev-list --tags --max-count=1`
          tag=`git describe --tags $revlist`

          major=`echo $tag | awk -F '.' '{ print $1 }'`
          minor=`echo $tag | awk -F '.' '{ print $2 }'`
          patch=`echo $tag | awk -F '.' '{ print $3 }'`
          patch=$(( $patch - $patch % 2 ))
          tag=$major.$minor.$patch

          git reset --hard
          git checkout $tag
          TAG=$tag make deploy-to-k8s-cluster
        '''.stripIndent())
      }
    }

    stage('Post') {
      steps {
        sh 'echo Posting'
      }
    }
  }

  post('Report') {
    always {
      echo "Anyway,finished the job."
     }
    // success {
    //   script {
    //     sh(script: 'bash $JENKINS_HOME/wechat-templates/send_wxmsg.sh successful')
    //  }
    //   script {
    //     // env.ForEmailPlugin = env.WORKSPACE
    //     emailext attachmentsPattern: 'TestResults\\*.trx',
    //     body: '${FILE,path="$JENKINS_HOME/email-templates/success_email_tmp.html"}',
    //     mimeType: 'text/html',
    //     subject: currentBuild.currentResult + " : " + env.JOB_NAME,
    //     to: '$DEFAULT_RECIPIENTS'
    //   }
    //  }
    // failure {
    //   script {
    //     sh(script: 'bash $JENKINS_HOME/wechat-templates/send_wxmsg.sh failure')
    //  }
    //   script {
    //     // env.ForEmailPlugin = env.WORKSPACE
    //     emailext attachmentsPattern: 'TestResults\\*.trx',
    //     body: '${FILE,path="$JENKINS_HOME/email-templates/fail_email_tmp.html"}',
    //     mimeType: 'text/html',
    //     subject: currentBuild.currentResult + " : " + env.JOB_NAME,
    //     to: '$DEFAULT_RECIPIENTS'
    //   }
    //  }
    // aborted {
    //   script {
    //     sh(script: 'bash $JENKINS_HOME/wechat-templates/send_wxmsg.sh aborted')
    //  }
    //   script {
    //     // env.ForEmailPlugin = env.WORKSPACE
    //     emailext attachmentsPattern: 'TestResults\\*.trx',
    //     body: '${FILE,path="$JENKINS_HOME/email-templates/fail_email_tmp.html"}',
    //     mimeType: 'text/html',
    //     subject: currentBuild.currentResult + " : " + env.JOB_NAME,
    //     to: '$DEFAULT_RECIPIENTS'
    //   }
    // }
  }
}
