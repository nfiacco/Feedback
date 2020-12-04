import { ThunkAction } from "redux-thunk";
import { sendRequest } from 'rpc/Ajax';
import { CheckSession } from 'rpc/Api';
import { RootAction, RootState } from "./model";

export function start(): ThunkAction<void, RootState, unknown, RootAction> {
  return async (dispatch: any) => {
    try {
      const checkSessionResponse = await sendRequest(CheckSession);
      dispatch({
        type: "login.authenticated",
        feedbackKey: checkSessionResponse.feedback_key,
      });
    } catch (e) {
    }

    dispatch({type: "done"});
  };
}