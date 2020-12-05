import { useParams } from "react-router-dom";

interface RouteParams {
  id: string;
}

export const Feedback: React.FC = () => {
  let { id } = useParams<RouteParams>();

  // check if the feedback key is valid, if not then redirect to not found page.
  // no need for redux state for this, just use local state
  return (
    <div>
      <h3>ID: {id}</h3>
    </div>
  );
}