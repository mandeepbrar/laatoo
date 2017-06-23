version="laatoo-0.1-1"
rm -fr tmp
mkdir -p tmp/$version
cp -R debian/* tmp/$version/
go build -o tmp/$version/usr/bin/laatoo laatoo/server
cd tmp
dpkg-deb --build $version
mv *.deb ..
