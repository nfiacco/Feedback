import { Dispatch } from "react";
import { useDispatch } from "react-redux";
import { combineReducers, createStore as createReduxStore } from "redux";
import { AppAction, appReducer, AppState } from "src/app/model";
import { LoginAction, loginReducer, LoginState } from "src/home/model";

export type RootAction = AppAction | LoginAction;

export interface RootState {
  app: AppState;
  login: LoginState;
}

export const useRootDispatch = () => useDispatch<Dispatch<RootAction>>();

export function createStore() {
  const rootReducer = combineReducers({app: appReducer, login: loginReducer});

  return createReduxStore(rootReducer);
}