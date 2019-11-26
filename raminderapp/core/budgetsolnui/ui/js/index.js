import React from 'react';
import './styles/app'
import {RequestBuilder, DataSource, EntityData} from 'uicommon';

function Initialize(appName, ins, mod, settings, def, req) {
  console.log("Initializing ui");
  _r("Methods", "Budget_Actions", Budget_Actions);
}

function publishBudget(form, submit, setData, dispatch) {
  return function(evt) {
    console.log(form.info.entityId)
    let req = RequestBuilder.URLParamsRequest({budgetId: form.info.entityId}, null);
    let res = DataSource.ExecuteService("publishbudget", req)
    res.then((res)=> {
      alert("Budget accounts published successfully");
      Window.redirect("/budgetentry")
    }, (err)=> {
      console.log("Could not publish budget", err);
      alert("Budget could not be published");
    })

    alert("published budget " + form.info.entityId)  
  }
}

const Expander = (props) => {
  console.log("props of expander ", props)
  let comps = []
  props.viewdepth? comps.push(<_uikit.Block className=" inlineblock " style={{width: props.viewdepth*15}}>&nbsp;</_uikit.Block>) : null
  comps.push(props.expanded? <_uikit.Icon className="fa fa-caret-down"/> :  <_uikit.Icon className="fa fa-caret-right"/>)
  return comps
}
  


function Budget_Actions(form, submit, reset, setData, dispatch) {
  console.log("my action buttons", dispatch)
  return (
    <_uikit.Block>
      <_uikit.ActionButton onClick={publishBudget(form, submit, setData, dispatch)} className="s10">Publish</_uikit.ActionButton>
      <_uikit.ActionButton onClick={submit()} className="submitBtn">Save</_uikit.ActionButton>
    </_uikit.Block>
  )
}

export {
  Initialize,
  Expander
}
