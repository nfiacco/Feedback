import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import { useSelector } from 'src/root/model';
import { Login } from '../login/Login';
import { Modal } from '../modal/Modal';
import styles from "./header.m.css";

export const Header: React.FC = () => {
  const [ showLoginModal, setShowLoginModal ] = useState(false);
  const isAuthenticated = useSelector(state => state.login.authenticated);
  const feedbackKey = useSelector(state => state.login.feedbackKey);

  return (
    <>
    <div className={styles.headerContainer}>
      <Link to={"/"}>Home</Link>
      <Link to={"/about"}>About</Link>
      {isAuthenticated ? (
        <>
          <div>You are logged in.</div>
          <Link to={"/feedback/" + feedbackKey}>My Feedback Link</Link>
        </>
      ): (
        <button type="button" onClick={()=>setShowLoginModal(true)}>Login</button>
      )}
    </div>
    <Modal show={showLoginModal} close={()=>setShowLoginModal(false)}>
        <Login/>
    </Modal>
  </>
  );
};