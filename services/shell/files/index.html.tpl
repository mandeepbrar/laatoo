<!doctype html>
<html>
<head>
  <meta charset="utf-8">
  <title>React Webpack Template Title</title>
  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
  <meta name="description" content="">
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no">
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/flexboxgrid/6.3.0/flexboxgrid.min.css" type="text/css" >
</head>
<body>
  <div id="app">APPLICATION CONTENT</div>

  <script type="text/javascript" src="/vendor.{{var "uiapplication"}}.js"></script>
  <script type="text/javascript" src="/app.{{var "uiapplication"}}.js"></script>
  <script type="text/javascript" src="/services.{{var "uiapplication"}}.conf.js"></script>
  <script type="text/javascript">
  <!--

    if (window.InitializeApplication) {
      console.log("Initializing app");
      window.InitializeApplication();
    }
  -->
  </script>

</body>
</html>
