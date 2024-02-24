# build and run server

set -ex
HERE=$(dirname $(realpath $BASH_SOURCE))
cd $HERE

cd bin
go build -o server.exe server.go
./server.exe