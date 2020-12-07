import React from "react";
import { Link } from "react-router-dom";

export const NotFound: React.FC = () => {
  return (
    <div>
      <Link to={"/"}>Home</Link>
      <h3>Not Found!</h3>
    </div>
  );
}