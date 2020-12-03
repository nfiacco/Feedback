import React, { useEffect } from 'react';
import GoogleLogin, { GoogleLoginResponseOffline } from 'react-google-login';
import { Provider, connect } from "react-redux";
import { applyMiddleware, createStore } from "redux";
import thunkMiddleware, { ThunkAction } from "redux-thunk";
import { GoogleLoginResponse } from "react-google-login"

const IS_PROD = process.env.NODE_ENV === "production";
const ROOT_DOMAIN = IS_PROD ? "https://api.anonymousfeedback.app" : "http://localhost:8080";

function isGoogleLoginResponse(response: GoogleLoginResponse | GoogleLoginResponseOffline): response is GoogleLoginResponse {
  return (response as GoogleLoginResponse).googleId !== undefined;
}

const responseGoogle = (response: GoogleLoginResponse | GoogleLoginResponseOffline) => {
  if (!isGoogleLoginResponse(response)) {
    return;
  }

  var id_token = response.getAuthResponse().id_token;

  var xhr = new XMLHttpRequest();
  xhr.open('POST', ROOT_DOMAIN + '/login');
  xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
  
  // enable sending cookies via CORS for development
  if (!IS_PROD) {
    xhr.withCredentials = true;
  }

  xhr.onload = function() {
    console.log('Signed in as: ' + xhr.responseText);
  };
  xhr.send(JSON.stringify({'idtoken': id_token}));
}

interface IState {
  loading: boolean;
}

type IAction = 
| {
  type: "loading"
}
| {
  type: "done"
};

function reducer(state: IState = INITIAL_STATE, action: IAction): IState {
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

const INITIAL_STATE: IState = {
    loading: true,
};

function createSignInStore() {
  const middlewares = [thunkMiddleware];
  const middlewareEnhancer = applyMiddleware(...middlewares)

  return createStore(reducer, INITIAL_STATE, middlewareEnhancer);
}

const Connector = connect(
  (state: IState) => ({
    loading: state.loading,
  }),
  {
    start: start,
  }
);

function sleep(ms: number) {
  return new Promise( resolve => setTimeout(resolve, ms) );
}

function start(): ThunkAction<void, IState, unknown, IAction> {
  return async (dispatch: any) => {
    await sleep(2000);
    dispatch({type: "done"});
  };
}

interface IProps {
  loading: boolean,
  start: () => void;
}

const SignInImpl: React.FC<IProps> = props => {
  useEffect(() => {
    props.start();
  });

  if (props.loading) {
    return <div>Loading!</div>
  }

  return (
    <GoogleLogin
      clientId="621422061156-f3f0o58fonsm9ohnqq5ngpa981c6k3hc.apps.googleusercontent.com"
      onSuccess={responseGoogle}
      isSignedIn={true}
    />
  )
}

const ConnectedSignIn = Connector(SignInImpl);

export function SignIn() {
  const store = createSignInStore();

  return (
    <Provider store={store}>
      <ConnectedSignIn/>
    </Provider>
  );
}
