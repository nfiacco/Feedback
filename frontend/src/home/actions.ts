import { RootAction, RootState } from "app/model";
import { GoogleLoginResponse, GoogleLoginResponseOffline } from "react-google-login";
import { ThunkAction } from "redux-thunk";
import { sendRequest } from 'rpc/Ajax';
import { Login } from "rpc/Api";

function isGoogleLoginResponse(response: GoogleLoginResponse | GoogleLoginResponseOffline): response is GoogleLoginResponse {
  return (response as GoogleLoginResponse).googleId !== undefined;
}

export function handleGoogleResponse(response: GoogleLoginResponse | GoogleLoginResponseOffline): ThunkAction<void, RootState, unknown, RootAction> {
  return async (dispatch: any) => {
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