MODULES_DIR=~mandeep/goprogs/bin/modules
cp $MODULES_DIR/filesystemstorage.tar.gz $MODULES_DIR/staticfileserver.tar.gz $MODULES_DIR/publicdirserver.tar.gz $MODULES_DIR/publicfilesserver.tar.gz \
  $MODULES_DIR/dataadapter.tar.gz $MODULES_DIR/mongodatabase.tar.gz $MODULES_DIR/shell.tar.gz dist/config/designer/modules
cp $MODULES_DIR/user.tar.gz $MODULES_DIR/role.tar.gz $MODULES_DIR/localauth.tar.gz $MODULES_DIR/dblogin.tar.gz dist/config/designer/modules
