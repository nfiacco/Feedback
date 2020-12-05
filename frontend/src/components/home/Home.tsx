import React from 'react';
import GoogleLogin from 'react-google-login';
import { Link } from 'react-router-dom';
import { useHandleGoogleResponse } from 'src/components/home/actions';
import { useSelector } from 'src/root/model';

const GOOGLE_CLIENT_ID = "621422061156-f3f0o58fonsm9ohnqq5ngpa981c6k3hc.apps.googleusercontent.com";

export const Home: React.FC = () => {
  const isAuthenticated = useSelector(state => state.login.authenticated);
  const feedbackKey = useSelector(state => state.login.feedbackKey);
  const handleGoogleResponse = useHandleGoogleResponse();

  return (
    <div>
      {isAuthenticated ? (
        <>
          <div>You are logged in.</div>
          <Link to={"/feedback/" + feedbackKey} >My Feedback Link</Link>
        </>
      ): (
        <GoogleLogin
          clientId={GOOGLE_CLIENT_ID}
          onSuccess={handleGoogleResponse}
        />
      )}
    </div>
  );
}