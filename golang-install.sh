wget https://dl.google.com/go/go1.16.15.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.16.15.linux-amd64.tar.gz
# shellcheck disable=SC2232
sudo export GOROOT=/usr/local/go
# shellcheck disable=SC2232
sudo export GOBIN="$GOROOT"/bin
# shellcheck disable=SC2232
sudo export GOPROXY=https://goproxy.cn,https://goproxy.io,direct
# shellcheck disable=SC2232
cd ~ || exit
mkdir go
# shellcheck disable=SC2232
sudo export GOPATH=~/go
# shellcheck disable=SC2232
sudo export PATH="$PATH":"$HOME"/bin:"$GOBIN":"$GOROOT"
echo go version
