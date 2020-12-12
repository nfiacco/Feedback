import { CircularProgress } from '@material-ui/core';
import React, { useState } from "react";
import ReactQuill from "react-quill";
import { useParams } from "react-router-dom";
import { useCheckKey, useSendFeedback } from "src/components/feedback/actions";
import { NotFound } from "src/components/notfound/NotFound";
import styles from "./feedback.m.css";
import "./quill.snow.css";

interface RouteParams {
  key: string;
}

export const Feedback: React.FC = () => {
  const { key } = useParams<RouteParams>();
  const { keyValid, loading } = useCheckKey(key);
  const [ inputText, setInputText ] = useState("");
  const { sendFeedback, sendStatus, clearSendStatus } = useSendFeedback(key, inputText);

  const sendMore = () => {
    setInputText("");
    clearSendStatus();
  }

  if (loading) {
    return <></>
  }

  if (sendStatus === "SENT") {
    return (
      <>
        <div>Sent!</div>
        <button type="button" onClick={sendMore}>Send more?</button>
      </>
    )
  }

  if (sendStatus === "ERROR") {
    return (
      <>
        <div>An Error Occurred!</div>
        {/* Don't clear the input text if an error happens, we want the user to have a chance to get their input back. */}
        <button type="button" onClick={clearSendStatus}>Try Again?</button>
      </>
    )
  }

  if (!keyValid) {
    return <NotFound/>
  }

  const sending = sendStatus === "SENDING";
  return (
    <div>
      <label className={styles.label}>
        Enter your feedback
      </label>
      <ReactQuill className={styles.textInput} theme="snow" value={inputText} onChange={setInputText}/>
      <button className={styles.submitButton} onClick={sending ? () => undefined : sendFeedback}>
        {sending ? <CircularProgress className={styles.submitSpinner}/> : "Send"}
      </button>
    </div>
  );
}