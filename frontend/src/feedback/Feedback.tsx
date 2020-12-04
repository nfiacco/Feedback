import { useParams } from "react-router-dom";

interface RouteParams {
  id: string;
}

export const Feedback: React.FC = () => {
  let { id } = useParams<RouteParams>();
  return (
    <div>
      <h3>ID: {id}</h3>
    </div>
  );
}