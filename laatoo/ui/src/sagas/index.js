import {viewSaga} from './View';
import {loginSaga} from './Security';
import {entitySaga} from './Entity';

export const Sagas = {
  ViewSaga: viewSaga,
  LoginSaga: loginSaga,
  EntitySaga: entitySaga
};

export function runSagas(sagaMiddleware, sagas) {
  sagas.map((x,i)=> {
    sagaMiddleware.run(x);
  })
}
