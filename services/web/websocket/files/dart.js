define(['dart_sdk'], function(dart_sdk) {
  'use strict';
  const core = dart_sdk.core;
  const dart = dart_sdk.dart;
  const dartx = dart_sdk.dartx;
  const _root = Object.create(null);
  const main = Object.create(_root);
  const src__test = Object.create(_root);
  main.myfunc = function() {
    core.print(src__test.testFunc());
    return "My string";
  };
  src__test.testFunc = function() {
    return "testfunc";
  };
  dart.trackLibraries("dart", {
    "main.dart": main,
    "src/test.dart": src__test
  }, null);
  // Exports:
  return {
    main: main,
    src__test: src__test
  };
});

//# sourceMappingURL=dart.js.map
