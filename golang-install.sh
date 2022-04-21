wget https://dl.google.com/go/go1.16.15.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.16.15.linux-amd64.tar.gz
# shellcheck disable=SC2232
sudo export GOBIN=/usr/local/go/bin
# shellcheck disable=SC2232
sudo export GOPROXY=https://goproxy.cn,https://goproxy.io,direct
# shellcheck disable=SC2232
sudo export GOPATH=~/go
# shellcheck disable=SC2232
sudo export PATH="$PATH":"$HOME"/bin:"$GOBIN"
echo go version
