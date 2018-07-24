CURR_DIR=`pwd`
cd binaryen
cmake .
make wasm-merge
cd $CURR_DIR
cd binaryenmerge
cmake .
make
cd $CURR_DIR
cp binaryenmerge/libbinaryenmerge.a binaryen/lib/libwasm.a binaryen/lib/libasmjs.a binaryen/lib/libemscripten-optimizer.a binaryen/lib/libpasses.a binaryen/lib/libir.a binaryen/lib/libcfg.a binaryen/lib/libsupport.a $CURR_DIR/lib
echo "Libraries built"

