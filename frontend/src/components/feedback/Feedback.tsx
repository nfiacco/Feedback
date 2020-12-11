import React, { useState } from "react";
import ReactQuill from "react-quill";
import "react-quill/dist/quill.snow.css";
import { Link, useParams } from "react-router-dom";
import { useCheckKey } from "src/components/feedback/actions";
import { NotFound } from "src/components/NotFound";
import styles from "./feedback.m.css";

interface RouteParams {
  key: string;
}

export const Feedback: React.FC = () => {
  const { key } = useParams<RouteParams>();
  const { keyValid, loading } = useCheckKey(key);
  const [ inputText, setInputText ] = useState("");

  // TODO: on submit, redirect to page "your feedback has been sent. write more?" so that they don't click multiple times
  // TODO: also add a debouncer maybe
  const onSubmit = () => {
    console.log("submitted: " + inputText);
  }

  if (loading) {
    return <></>
  }

  if (!keyValid) {
    return <NotFound/>
  }

  return (
    <div>
      <Link to={"/"}>Home</Link>
      <label className={styles.label}>
        Enter your feedback
      </label>
      <ReactQuill className={styles.textInput} theme="snow" value={inputText} onChange={setInputText}/>
      <button className={styles.submitButton} onClick={onSubmit}>Send</button>
    </div>
  );
}