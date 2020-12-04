import { handleGoogleResponse } from 'home/actions';
import React from 'react';
import GoogleLogin from 'react-google-login';

export const Home: React.FC = props => {
  return (
    <GoogleLogin
      clientId="621422061156-f3f0o58fonsm9ohnqq5ngpa981c6k3hc.apps.googleusercontent.com"
      onSuccess={handleGoogleResponse}
      isSignedIn={true}
    />
  )
}
