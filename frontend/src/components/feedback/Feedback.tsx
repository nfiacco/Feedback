import React from "react";
import { Link, useParams } from "react-router-dom";
import { useCheckKey } from "src/components/feedback/actions";
import { NotFound } from "src/components/NotFound";

interface RouteParams {
  key: string;
}

export const Feedback: React.FC = () => {
  const { key } = useParams<RouteParams>();
  const { keyValid, loading } = useCheckKey(key);

  if (loading) {
    return <></>
  }

  if (!keyValid) {
    return <NotFound/>
  }

  return (
    <div>
      <Link to={"/"} >Home</Link>
      <h3>Key: {key}</h3>
    </div>
  );
}