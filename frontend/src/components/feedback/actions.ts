import { useEffect, useState } from "react";
import { sendRequest } from "src/rpc/ajax";
import { CheckKey } from "src/rpc/api";

export const useCheckKey = (key: string): {keyValid: boolean, loading: boolean} => {
  const [loading, setLoading] = useState(true);
  const [keyValid, setKeyValid] = useState(false);

  useEffect(() => {
    let isCancelled = false;

    const checkKey = async () => {
      let keyValid;
      try {
        await sendRequest(CheckKey, { key });
        keyValid = true;
      } catch(e) {
        keyValid = false;
      }

      if (!isCancelled) {
        setKeyValid(keyValid);
        setLoading(false);
      }
    }
    checkKey();

    return () => { isCancelled = true };
  }, [key]);

  return { keyValid, loading }
}