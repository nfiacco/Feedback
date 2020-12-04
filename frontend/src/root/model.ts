import { applyMiddleware, combineReducers, createStore as createReduxStore } from "redux";
import thunkMiddleware, { ThunkAction } from "redux-thunk";
import { AppAction, appReducer, AppState } from "src/app/model";
import { LoginAction, loginReducer, LoginState } from "src/home/model";

export type RootAction = AppAction | LoginAction;

export interface RootState {
  app: AppState;
  login: LoginState;
}

export type AsyncAction = ThunkAction<void, RootState, unknown, RootAction>;

export function createStore() {
  const middlewares = [thunkMiddleware];
  const middlewareEnhancer = applyMiddleware(...middlewares)
  const rootReducer = combineReducers({app: appReducer, login: loginReducer});

  return createReduxStore(rootReducer, middlewareEnhancer);
}