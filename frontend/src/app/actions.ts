import { AppAction, AppState } from "app/model";
import { ThunkAction } from "redux-thunk";
import { sendRequest } from 'rpc/Ajax';
import { CheckSession } from 'rpc/Api';

export function start(): ThunkAction<void, AppState, unknown, AppAction> {
  return async (dispatch: any) => {
    const isAuthenticated = await checkSession();
    if (isAuthenticated) {
      dispatch({type: "login.authenticated"});
    }
    dispatch({type: "done"});
  };
}

async function checkSession(): Promise<boolean> {
  try {
    await sendRequest(CheckSession);
    return true;
  } catch (e) {
    return false;
  }
}
