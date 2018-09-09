#!/bin/sh

git fetch --all

git reset --hard origin/master

git pull

kill -9 $(pidof apiserver)

cd lq186.com/apiserver

go build

nohup ./apiserver >> apiserver.out 2>&1 &
