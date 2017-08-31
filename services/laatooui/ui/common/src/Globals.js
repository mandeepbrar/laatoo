let native = typeof(document) == 'undefined'
let storage = native?  {}: localStorage
let application = native?  {}: document.InitConfig
let wind = native? {}: window
application.native = native

module.exports = {
  Storage: storage,
  Application: application,
  Window: wind
}
