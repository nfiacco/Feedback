import styles from "./about.m.css";
import image from "./elmo.gif";

export const About: React.FC = () => {
  return (
    <>
      <div className={styles.info}>
        I made this for fun.
      </div>
      <div className={styles.body}>
        <img src={image} alt="loading..." />
      </div>
    </>
  );
}