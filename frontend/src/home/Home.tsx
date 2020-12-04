import React from 'react';
import GoogleLogin, { GoogleLoginResponse, GoogleLoginResponseOffline } from 'react-google-login';
import { connect } from 'react-redux';
import { handleGoogleResponse } from 'src/home/actions';
import { RootState } from 'src/root/model';

interface HomeProps {
  isAuthenticated: boolean;
  handleGoogleResponse: (response: GoogleLoginResponse | GoogleLoginResponseOffline) => void;
}

const HomeImpl: React.FC<HomeProps> = props => {
  if (props.isAuthenticated) {
    return <div>You are logged in.</div>
  }

  return (
    <GoogleLogin
      clientId="621422061156-f3f0o58fonsm9ohnqq5ngpa981c6k3hc.apps.googleusercontent.com"
      onSuccess={props.handleGoogleResponse}
      isSignedIn={true} // this will automatically trigger onSuccess if the user is already logged-in via Google
    />
  )
}

const Connector = connect(
  (state: RootState) => ({
    isAuthenticated: state.login.authenticated,
  }),
  {
    handleGoogleResponse: handleGoogleResponse,
  }
);

export const Home = Connector(HomeImpl);