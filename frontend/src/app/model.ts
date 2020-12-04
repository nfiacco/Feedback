import { LoginAction, loginReducer, LoginState } from "home/model";
import { applyMiddleware, combineReducers, createStore as createReduxStore } from "redux";
import thunkMiddleware from "redux-thunk";

export type AppAction = 
| {
  type: "loading"
}
| {
  type: "done"
};

const INITIAL_APP_STATE: AppState = {
  loading: true,
}

export interface AppState {
  loading: boolean;
}

function appReducer(state: AppState = INITIAL_APP_STATE, action: AppAction): AppState {
  switch (action.type) {
    case "loading":
    return {
      ...state,
      loading: true,
    }
    case "done":
    return {
      ...state,
      loading: false,
    }
    default:
    return state;
  }
}

export type RootAction = AppAction | LoginAction;

export interface RootState {
  app: AppState;
  login: LoginState;
}

export function createStore() {
  const middlewares = [thunkMiddleware];
  const middlewareEnhancer = applyMiddleware(...middlewares)
  const rootReducer = combineReducers({app: appReducer, login: loginReducer});

  return createReduxStore(rootReducer, middlewareEnhancer);
}