import { AsyncAction } from "src/root/model";
import { sendRequest } from 'src/rpc/Ajax';
import { CheckSession } from 'src/rpc/Api';

export function start(): AsyncAction {
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