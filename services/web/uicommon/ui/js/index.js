import { DataSource, RequestBuilder, EntityData } from './sources/DataSource';
import {Response} from './sources/ResponseCodes';
import {createAction, formatUrl, LaatooError, hasPermission} from './utils';
import Colors from './colors'
import { RestDataSource } from './http/RestSource';


function Initialize(appName, ins, mod, settings, def, req) {
  console.log("****************Init uicommon*****")
  _r("DataSourceHandlers", "http", new RestDataSource())
  if(settings && settings.EntityPrefix) {
    EntityData.SetPrefix(settings.EntityPrefix)
  }
}


export {
  Initialize,
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
