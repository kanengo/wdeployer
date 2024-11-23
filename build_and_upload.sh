#!/bin/sh

if [ $# -eq 0 ]; then
  echo "没有提供任何参数"
  exit  1
else
  echo "上传版本: $1"
fi

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o kube-mono_linux_amd64 ./cmd/kube-mono \
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o kube-mono_darwin_amd64 ./cmd/kube-mono \
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o kube-mono_win_amd64.exe ./cmd/kube-mono \
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o kube-mono_darwin_arm64 ./cmd/kube-mono

user=packageadmin
token=2d045b553a46d7985ec689ad58679e3472cb66bc
url=https://storehouse.dev.nextlove.online
team=backend
package=kube-mono

for bin in "kube-mono_darwin_amd64" "kube-mono_darwin_arm64" "kube-mono_linux_amd64" "kube-mono_win_amd64.exe"; do
curl --user "$packageadmin":"$token" \
     --upload-file "$bin" \
     "$url"/api/packages/"$team"/generic/"$package"/"$1"/"$bin"
done

