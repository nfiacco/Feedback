import styles from "./about.m.css";

export const About: React.FC = () => {
  return (
    <>
      <div className={styles.info}>
        About Us
      </div>
      <div className={styles.body}>
        We created Anonymous Feedback based on the belief that feedback is a gift, and critical to our personal development.
        However, the reality is that giving and receiving feedback can be pretty awkward. By providing a mechanism
        to give and receive anonymous feedback, we hope to make constant feedback a cultural norm.
        <br/><br/>
        We hope you enjoy the app, and leave us a review using *our* feedback link below:
        <br/><br/>
        TODO
      </div>
    </>
  );
}