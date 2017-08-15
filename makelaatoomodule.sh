#!/bin/sh

rm -fr tmp/$1
mkdir -p tmp/$1/objects
go build -buildmode=plugin -o tmp/$1/objects/$1.so $2/*.go
if [ -e $2/config ]
 then cp -R $2/config/* tmp/$1/
fi
if [ -e $2/files ]
 then cp -R $2/files tmp/$1/
fi
tar -czvf tmp/$1.tar.gz -C tmp/ $1
mkdir -p $3
cp tmp/$1.tar.gz  $3
