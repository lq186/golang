#!/bin/sh

kill -9 $(pidof apiserver)

cd lq186.com/apiserver

go build

nohup ./apiserver >> apiserver.out 2>&1 &
