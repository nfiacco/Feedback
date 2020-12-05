import { useRootDispatch } from 'src/root/model';
import { sendRequest } from 'src/rpc/Ajax';
import { CheckSession } from 'src/rpc/Api';

export function useStart() {
  const dispatch = useRootDispatch();
  return async () => {
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