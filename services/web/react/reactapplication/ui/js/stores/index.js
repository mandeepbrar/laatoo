const redux = require('redux');
//import { createAction, createStore, Sagas} from 'uicommon'
//import {Errors} from '../messages'
import createSagaMiddleware from 'redux-saga';
import {all, call, fork, spawn} from 'redux-saga/effects';
import '../reducers/Interactions'
import '../reducers/Messages'

function runSagas(sagaMiddleware, sagas) {
  if(sagas) {
    Object.keys(sagas).forEach(function(sagaId){
      let saga = sagas[sagaId];
      sagaMiddleware.run(saga);
    });
  }
}


function* rootSaga () {
  let sagas = Application.AllRegItems("Sagas")  
  console.log("running sagas", sagas)
  /*yield all(Object.entries(sagas).map(function(name, saga) {
    spawn(function*() {
      console.log("saga", name)
      while (true) {
        try {
          console.log("while loop ", name)
          yield call(saga)
          break
        } catch (e) {
          console.log(e)
        }
      }
    })
    console.log("running saga", name)
  }))*/
  /*for (let saga of Object.values(sagas)) {
    console.log("saga", saga)
    spawn(function*() {
      console.log("saga", saga)
      while (true) {
        try {
          console.log("while loop ", saga)
          yield call(saga)
          break
        } catch (e) {
          console.log(e)
        }
      }
    })
  }*/
/*  while (true) {
    try {
      console.log("while loop ")
      let arr = Object.values(sagas).map(saga=>fork(saga))
      yield arr;
    }catch (e) {
      console.log(e)
    }
  }*/
  console.log("exitting saga")
  
}


function configureStore() {
  let reducers = Application.AllRegItems("Reducers")
  if(!reducers) {
    reducers = {};
  }
  console.log("reducers in store", reducers);
  let middleware = [];
  let enhancers = [];


  const sagaMiddleware = createSagaMiddleware();
  //enhancers = redux.compose(redux.applyMiddleware(sagaMiddleware, ...middleware), ...enhancers);

  // mount it on the Store
  const store = redux.createStore(redux.combineReducers(reducers), redux.applyMiddleware(sagaMiddleware, ...middleware), ...enhancers);
  console.log("created store", store)

  // then run the saga

  for (let [sagaName, saga] of Object.entries(Application.AllRegItems("Sagas"))) {
    try {
      console.log("running saga", sagaName)
      sagaMiddleware.run(saga)
    }catch(ex) {
      console.log("Couldnt run saga", sagaName)
    }
  }
//  sagaMiddleware.run(rootSaga);
  console.log("run sagas complete", Application.AllRegItems("Sagas"))
  return store;
}


export default configureStore

/*
module.exports = function(reducers, initialState, middleware, sagas, enhancers) {
  const store = createStore(reducers, initialState, middleware, [Sagas.LoginSaga, Sagas.ViewSaga, Sagas.GroupLoadSaga, Sagas.EntitySaga, MehfilSaga, PostSaga, MediaSaga, CommentsSaga, ArticleSaga, MiscSaga, ...sagas], enhancers);
  return store
}


const redux = require('redux');
//const reducers = require('../reducers');
import {history, createAction} from 'laatoo';
import {createStore, Errors, Reducers, MehfilActions} from 'mehfiluicommon'
import {Reducer} from 'redux-director';

function devtools(middleware, enhancers) {
  if(window.devToolsExtension) {
    enhancers.push(window.devToolsExtension());
  }
  function logger({ getState }) {
    return (next) => (action) => {
      console.log('will dispatch1', action)

      // Call the next dispatch method in the middleware chain.
      let returnValue = next(action)

      console.log('state after dispatch', getState())

      // This will likely be the action itself, unless
      // a middleware further in chain changed it.
      return returnValue
    }
  }
  middleware.push(logger);
}

module.exports = function(initialState) {
  let middleware = [];
  let enhancers = [];
  devtools(middleware, enhancers);
  let reducers = Object.assign({}, Reducers, {router: Reducer})

  const store = createStore(reducers, initialState, middleware, [], enhancers);
  if (module.hot) {
    // Enable Webpack hot module replacement for reducers
    module.hot.accept('../reducers', () => {
      const nextReducer = require('../reducers')
      store.replaceReducer(nextReducer)
    })
  }
  return store
}*/
