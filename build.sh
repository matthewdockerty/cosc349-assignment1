#!/bin/bash
/usr/local/go/bin/go build -o /vagrant/webapp/server /vagrant/webapp/server.go /vagrant/webapp/database.go

if [ $? == 0 ]; then
    echo "Build successful"
else
    echo "Build failed"
    exit 1
fi
