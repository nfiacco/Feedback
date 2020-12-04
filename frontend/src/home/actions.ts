import { GoogleLoginResponse, GoogleLoginResponseOffline } from "react-google-login";
import { sendRequest } from 'rpc/Ajax';
import { Login } from "rpc/Api";

function isGoogleLoginResponse(response: GoogleLoginResponse | GoogleLoginResponseOffline): response is GoogleLoginResponse {
  return (response as GoogleLoginResponse).googleId !== undefined;
}

export const handleGoogleResponse = async (response: GoogleLoginResponse | GoogleLoginResponseOffline) => {
  if (!isGoogleLoginResponse(response)) {
    return;
  }

  const id_token = response.getAuthResponse().id_token;
  const payload = {'idtoken': id_token};
  const loginResponse = await sendRequest(Login, payload);
  console.log(JSON.stringify(loginResponse));
}