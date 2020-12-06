import { GoogleLoginResponse, GoogleLoginResponseOffline } from "react-google-login";
import { useDispatch } from "src/root/model";
import { sendRequest } from 'src/rpc/ajax';
import { Login } from "src/rpc/api";

function isGoogleLoginResponse(response: GoogleLoginResponse | GoogleLoginResponseOffline): response is GoogleLoginResponse {
  return (response as GoogleLoginResponse).googleId !== undefined;
}

export type GoogleLoginHandler = (response: GoogleLoginResponse | GoogleLoginResponseOffline) => void;

export function useHandleGoogleResponse(): GoogleLoginHandler {
  const dispatch = useDispatch();
  return async (response: GoogleLoginResponse | GoogleLoginResponseOffline) => {
    if (!isGoogleLoginResponse(response)) {
      return;
    }

    const id_token = response.getAuthResponse().id_token;
    const payload = {'id_token': id_token};
    const loginResponse = await sendRequest(Login, payload);
    dispatch({
      type: "login.authenticated",
      feedbackKey: loginResponse.feedback_key,
    });
  };
}