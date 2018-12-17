#!/usr/bin/env bash

PACKAGE="mlbcli"
VERSION="0.1.0"
#PLATFORMS="linux/386,windows/386,darwin/amd64"
PLATFORMS="darwin/amd64"

go get github.com/tidwall/match
go get github.com/apcera/termtables
go get github.com/tidwall/gjson
go get gopkg.in/jarcoal/httpmock.v1

go test
if [ "$?" -ne 0 ]; then
  exit 14
fi

for PLATFORM in $(echo ${PLATFORMS}|sed 's/,/ /g'); do
  GOOS="$(echo ${PLATFORM}|cut -d/ -f1)"
  GOARCH="$(echo ${PLATFORM}|cut -d/ -f2)"
  OUTPUT="${PACKAGE}-${VERSION}-${GOOS}-${GOARCH}"
  if [ "$(echo $GOOS|grep -i windows)" ]; then
      OUTPUT="${OUTPUT}.exe"
  fi
  rm -f $OUTPUT ${OUTPUT}.zip
  env GOOS=$GOOS GOARCH=$GOARCH go build -o $OUTPUT
  # zip -q -9 ${OUTPUT}.zip $OUTPUT
  shasum -a  256 ${OUTPUT}
done
