const redux = require('redux');
import {Application, createAction, createStore, Sagas, history, Window} from 'reactuibase'
//import {Errors} from '../messages'
import createSagaMiddleware from 'redux-saga';

function runSagas(sagaMiddleware, sagas) {
  sagas.map((x,i)=> {
    sagaMiddleware.run(x);
  })
}

function configureStore() {
  let reducers = Application.Reducers
  if(!reducers) {
    reducers = {};
  }

  let middleware = [];
  let enhancers = [];


  const sagaMiddleware = createSagaMiddleware();
  enhancers = redux.compose(redux.applyMiddleware(sagaMiddleware, ...middleware), ...enhancers);


  // mount it on the Store
  const store = redux.createStore(redux.combineReducers(reducers), {}, enhancers);

  // then run the saga
  runSagas(sagaMiddleware, Application.Sagas);
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
