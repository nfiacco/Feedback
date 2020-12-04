import { RootState } from 'app/model';
import { handleGoogleResponse } from 'home/actions';
import React from 'react';
import GoogleLogin, { GoogleLoginResponse, GoogleLoginResponseOffline } from 'react-google-login';
import { connect } from 'react-redux';

interface HomeProps {
  handleGoogleResponse: (response: GoogleLoginResponse | GoogleLoginResponseOffline) => void;
}

const HomeImpl: React.FC<HomeProps> = props => {
  return (
    <GoogleLogin
      clientId="621422061156-f3f0o58fonsm9ohnqq5ngpa981c6k3hc.apps.googleusercontent.com"
      onSuccess={props.handleGoogleResponse}
      isSignedIn={true}
    />
  )
}

const Connector = connect(
  (state: RootState) => ({}),
  {
    handleGoogleResponse: handleGoogleResponse,
  }
);

export const Home = Connector(HomeImpl);