import { CircularProgress } from "@material-ui/core";
import classNames from "classnames";
import React, { useState } from "react";
import GoogleLogin, { GoogleLoginResponse, GoogleLoginResponseOffline } from "react-google-login";
import isEmail from 'validator/lib/isEmail';
import { useHandleGoogleResponse } from "./actions";
import styles from "./login.m.css";

const GOOGLE_CLIENT_ID = "621422061156-f3f0o58fonsm9ohnqq5ngpa981c6k3hc.apps.googleusercontent.com";

interface LoginProps {
  closeModal: () => void;
}

export const Login: React.FC<LoginProps> = props => {
  const [ loading, setLoading ] = useState(false);
  const handleGoogleResponse = useHandleGoogleResponse();
  const onSuccess = (response: GoogleLoginResponse | GoogleLoginResponseOffline) => {
    handleGoogleResponse(response);
    props.closeModal();
  }

  if (loading) {
    return (
      <div className={styles.loadingContainer}>
        <CircularProgress className={styles.spinner}/>
      </div>
    )
  }

  return (
    <div className={styles.loginContainer}>
      <h2 className={classNames(styles.center, styles.header)}>Anonymous Feedback</h2>
      <div className={classNames(styles.center, styles.loginGroup)}>
        <GoogleLogin
          clientId={GOOGLE_CLIENT_ID}
          onSuccess={onSuccess}
          render={renderProps => <GoogleButton onClick={renderProps.onClick} setLoading={setLoading}/>}
          onFailure={()=>setLoading(false)}
        />
        <div className={styles.marginTop}>
          or
        </div>
        <EmailLogin/>
      </div>
    </div>
  );
};

interface GoogleButtonProps {
  onClick: () => void;
  setLoading: (loading: boolean) => void;
}

const GoogleButton: React.FC<GoogleButtonProps> = props => {
  const onClick = () => {
    props.setLoading(true);
    props.onClick();
  }

  return (
    <button className={styles.loginButton} onClick={onClick}>
      <div className={styles.buttonContentWrapper}>
        <span className={styles.iconWrapper}>
          <svg width="18" height="18" xmlns="http://www.w3.org/2000/svg">
            <g fill="#000" fillRule="evenodd">
              <path d="M9 3.48c1.69 0 2.83.73 3.48 1.34l2.54-2.48C13.46.89 11.43 0 9 0 5.48 0 2.44 2.02.96 4.96l2.91 2.26C4.6 5.05 6.62 3.48 9 3.48z" fill="#EA4335"/>
              <path d="M17.64 9.2c0-.74-.06-1.28-.19-1.84H9v3.34h4.96c-.1.83-.64 2.08-1.84 2.92l2.84 2.2c1.7-1.57 2.68-3.88 2.68-6.62z" fill="#4285F4"></path><path d="M3.88 10.78A5.54 5.54 0 0 1 3.58 9c0-.62.11-1.22.29-1.78L.96 4.96A9.008 9.008 0 0 0 0 9c0 1.45.35 2.82.96 4.04l2.92-2.26z" fill="#FBBC05"></path><path d="M9 18c2.43 0 4.47-.8 5.96-2.18l-2.84-2.2c-.76.53-1.78.9-3.12.9-2.38 0-4.4-1.57-5.12-3.74L.97 13.04C2.45 15.98 5.48 18 9 18z" fill="#34A853"/>
              <path fill="none" d="M0 0h18v18H0z"></path>
            </g>
          </svg>
        </span>
        Sign in with Google
      </div>
    </button>
  )
}

const EmailLogin: React.FC = () => {
  const [ email, setEmail ] = useState("");
  const [ isValid, setIsValid ] = useState(true);

  const onKeydown = (event: React.KeyboardEvent<HTMLInputElement>) => {
    event.stopPropagation();
    if(event.key === "Escape") {
      event.currentTarget.blur();
    }
  }

  const validateEmail = () => {
    setIsValid(isEmail(email));
  }

  const onSubmit = () => {
    validateEmail();
    console.log(email);
  }

  let classes = [styles.emailInput]
  if (!isValid) {
    classes.push(styles.invalidEmailBorder);
  }

  return (
    <div className={styles.marginTop}>
      <input
        type="text"
        id="email"
        name="email"
        autoComplete="email"
        placeholder="Email"
        className={classNames(classes)}
        onKeyDown={onKeydown}
        onChange={e => setEmail(e.target.value)}
        onBlur={validateEmail}
      />
      {!isValid && <div className={styles.invalidEmailLabel}>Please check your email.</div>}
      <input type="submit" value="Continue" className={styles.emailSubmit} onClick={onSubmit}/>
    </div>
  )
};