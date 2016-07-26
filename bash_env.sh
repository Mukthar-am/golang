export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin

export GOWORKSPACE=`pwd`
export GOPATH=$GOWORKSPACE
export GOBIN=$GOPATH/bin/



# run build
go build src/helloworld/maloadtouch/
go install src/helloworld/maloadtouch

go install src/helloworld/src/Touchload.go
go install src/helloworld/src/Counter.go
