import {viewSaga} from './View';
import {loginSaga} from './Security';
import {entitySaga} from './Entity';
import {groupLoadSaga} from './GroupLoad';

export const Sagas = {
  ViewSaga: viewSaga,
  LoginSaga: loginSaga,
  GroupLoadSaga: groupLoadSaga,
  EntitySaga: entitySaga
};
