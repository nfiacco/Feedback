import { GoogleLoginResponse, GoogleLoginResponseOffline } from "react-google-login";
import { useRootDispatch } from "src/root/model";
import { sendRequest } from 'src/rpc/Ajax';
import { Login } from "src/rpc/Api";

function isGoogleLoginResponse(response: GoogleLoginResponse | GoogleLoginResponseOffline): response is GoogleLoginResponse {
  return (response as GoogleLoginResponse).googleId !== undefined;
}

export type GoogleLoginHandler = (response: GoogleLoginResponse | GoogleLoginResponseOffline) => void;

export function useHandleGoogleResponse(): GoogleLoginHandler {
  const dispatch = useRootDispatch();
  return async (response: GoogleLoginResponse | GoogleLoginResponseOffline) => {
    if (!isGoogleLoginResponse(response)) {
      return;
    }

    const id_token = response.getAuthResponse().id_token;
    const payload = {'idtoken': id_token};
    const loginResponse = await sendRequest(Login, payload);
    dispatch({
      type: "login.authenticated",
      feedbackKey: loginResponse.feedback_key,
    });
  };
}