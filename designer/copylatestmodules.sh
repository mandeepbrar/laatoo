LAATOO_MODULES_DIR=/home/mandeep/goprogs/src/laatoo/modules
DESIGNER_MODULES_DIR=/home/mandeep/goprogs/src/laatoo/designer/modules
cp $LAATOO_MODULES_DIR/filesystemstorage.tar.gz $LAATOO_MODULES_DIR/staticfileserver.tar.gz $LAATOO_MODULES_DIR/publicdirserver.tar.gz $LAATOO_MODULES_DIR/publicfilesserver.tar.gz \
  $LAATOO_MODULES_DIR/dataadapter.tar.gz $LAATOO_MODULES_DIR/mongodatabase.tar.gz $LAATOO_MODULES_DIR/shell.tar.gz $LAATOO_MODULES_DIR/ui.tar.gz $LAATOO_MODULES_DIR/reactuibase.tar.gz dist/config/designer/modules
cp $LAATOO_MODULES_DIR/user.tar.gz $LAATOO_MODULES_DIR/role.tar.gz $LAATOO_MODULES_DIR/localauth.tar.gz $LAATOO_MODULES_DIR/dblogin.tar.gz dist/config/designer/modules
cp $DESIGNER_MODULES_DIR/designerui.tar.gz dist/config/designer/modules
