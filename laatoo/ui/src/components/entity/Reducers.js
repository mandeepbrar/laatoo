import  {EntityFormReducer} from '../../reducers/EntityForm';
import  {EntityReducer} from '../../reducers/Entity';
import  {ViewReducer} from '../../reducers/View';
import { combineReducers } from 'redux';

export function getCreateReducer(name) {
  let reducerName = name.toUpperCase()+"_Form";
  let reducers = {};
  reducers[reducerName] = EntityFormReducer(reducerName);
  let createReducer = combineReducers(reducers);
  return createReducer;
}

export function getUpdateReducer(name) {
  let reducerName = name.toUpperCase()+"_Form";
  let reducers = {};
  reducers[reducerName] = EntityFormReducer(reducerName);
  let updateReducer = combineReducers(reducers);
  return updateReducer;
}

export function getViewReducer(name) {
  let reducerName = name.toUpperCase()+"_View";
  let reducers = {}
  reducers[reducerName] = ViewReducer(reducerName);
  let viewReducer = combineReducers(reducers);
  return viewReducer;
}

export function getDisplayReducer(name) {
  let reducerName = name.toUpperCase()+"_Display";
  let reducers = {}
  reducers[reducerName] = EntityReducer(reducerName);
  let displayReducer = combineReducers(reducers);
  return displayReducer;
}
