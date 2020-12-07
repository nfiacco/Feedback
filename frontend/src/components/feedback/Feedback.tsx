import React, { useState } from "react";
import { Link, useParams } from "react-router-dom";
import { useCheckKey } from "src/components/feedback/actions";
import { NotFound } from "src/components/NotFound";

interface RouteParams {
  key: string;
}

export const Feedback: React.FC = () => {
  const { key } = useParams<RouteParams>();
  const { keyValid, loading } = useCheckKey(key);
  const [ inputText, setInputText ] = useState("");

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
      <label>
        Enter your feedback
      </label>
      <input type="text" value={inputText} onChange={e => setInputText(e.target.value)} />
      <button type="button" onClick={onSubmit}>Submit</button>
    </div>
  );
}