
export class LaatooError extends Error {
  constructor(type, rootError, args) {
    super(type);
    this.name = this.constructor.name;
    this.message = type;
    if (typeof Error.captureStackTrace === 'function') {
      Error.captureStackTrace(this, this.constructor);
    } else {
      this.stack = (new Error(type)).stack;
    }
    this.type = type;
    this.rootError = rootError;
    this.args = args;
  }
}


export function createAction(type, payload, meta) {
  let error =  payload instanceof Error;
  console.log("created action", type, payload, meta, error)
  return {
    type,
    payload,
    meta,
    error
  }
}

export function formatUrl(url, params) {
  var newurl = url;
  if(params) {
    for (var key in params) {
        let val = params[key];
        newurl = newurl.replace(new RegExp(":"+key, "g"), val);
    }
  }
  return newurl
}

export function hasPermission(permission) {
  let hasPermission = true;
  if(permission && permission!="") {
    let permissions = localStorage.permissions;
    if(permissions) {
      if(permissions.indexOf(permission) < 0) {
        hasPermission = false;
      }
    }
  }
  return hasPermission;
}
