import React from 'react';
import GoogleLogin from 'react-google-login';
import { useSelector } from 'react-redux';
import { RootState } from 'src/root/model';
import { useHandleGoogleResponse } from './actions';

export const Home: React.FC = () => {
  const isAuthenticated = useSelector((state: RootState) => state.login.authenticated);
  const handleGoogleResponse = useHandleGoogleResponse();

  if (isAuthenticated) {
    return <div>You are logged in.</div>
  }

  return (
    <GoogleLogin
      clientId="621422061156-f3f0o58fonsm9ohnqq5ngpa981c6k3hc.apps.googleusercontent.com"
      onSuccess={handleGoogleResponse}
    />
  )
}