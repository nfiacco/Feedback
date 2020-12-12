import React from 'react';
import styles from './home.m.css';

export const Home: React.FC = () => {
  return (
    <>
    <div className={styles.title}>
      Welcome to Anonymous Feedback!
    </div>
    <div className={styles.body}>
      Login to get started.
    </div>
    </>
  );
}