import { Response, DataSource, RequestBuilder, EntityData } from './sources/DataSource';
import {createAction, formatUrl, LaatooError, hasPermission} from './utils';
import Color from './colors'


function Initialize(appName, ins, mod, settings, def, req) {
  if(settings && settings.EntityPrefix) {
    EntityData.SetPrefix(settings.EntityPrefix)
  }
}


export {
  RequestBuilder,
  DataSource,
  Response,
  EntityData,
  Colors,
  formatUrl,
  createAction,
  LaatooError,
  hasPermission
}
