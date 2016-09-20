//renderFunc = function (panelname, contentparams, styleattribs )

if(window.PanelsRegistry != null) {
  window.PanelsRegistry = new Array()
}

class PanelsRegistry {
  RegisterPanel(panelType, renderFunc) {
    window.PanelsRegistry[panelType] = renderFunc
  }
  GetPanel(panelType) {
    return window.PanelsRegistry[panelType]
  }
}

export default PanelsRegistry
