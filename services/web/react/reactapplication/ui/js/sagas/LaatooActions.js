import  Actions from '../actions';
import {hasPermission } from 'uicommon';
import { takeEvery, call, put, fork } from 'redux-saga/effects'
import { createAction} from 'uicommon'

var initialState = {
};

function* actionFunc (laatooaction, params){
    console.log("executing action in saga", laatooaction, params)
    if(laatooaction.permission) {
        if (!hasPermission(laatooaction.permission)) {
            return
        }
    }
    switch(laatooaction.actiontype) {
        case "method":
            let actionMethod = laatooaction.method
            if(actionMethod && typeof(actionMethod) === "string") {
                actionMethod = _reg("Methods", laatooaction.method)
            }
            if(actionMethod && (typeof(actionMethod) === "function")) {
                actionMethod(params);
            }      
            return;
        case "openinteraction":
            let comp = Window.resolvePanel("block", laatooaction.blockid, params)
            yield put(createAction(Actions.SHOW_INTERACTION_COMP, Object.assign(laatooaction, {component: comp}), null))
            return;
        case "closeinteraction":
            yield put(createAction(Actions.CLOSE_INTERACTION_COMP, action, null))
            //fall over to default case and check for URL 
        default:
            if(laatooaction.url) {
                Window.redirect(laatooaction.url, params, laatooaction.newWindow);
            }
            if(laatooaction.page) {
                Window.redirectPage(laatooaction.page, params);
            }
            return;
    }
}


function* LaatooActions(action) {
    try {
        console.log("got laatoo action in saga", action)
        if(action.meta) {
            let laatooActionToExec = action.meta
            if(action.meta.Id) {
                laatooActionToExec = _reg('Actions', props.Id)
            }
            let params = action.payload
            yield actionFunc(laatooActionToExec, params)
        }
        if(action.meta.successCallback) {
            action.meta.successCallback({resp: resp, payload: action.payload})
        }
    } catch (e) {
      if(action.meta.failureCallback) {
        action.meta.failureCallback(e)
      } else {
        if(Window.handleError) {
          Window.handleError(e)
        }
      }
    }
  }

function* laatooActionSaga() {
    yield takeEvery(Actions.LAATOO_ACTION, LaatooActions)
}

function* actionsSaga() {
    yield fork(laatooActionSaga)
}
  
Application.Register('Sagas', "laatooActionsSaga", actionsSaga)