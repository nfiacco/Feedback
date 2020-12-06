import { useEffect, useState } from "react";
import { sendRequest } from "src/rpc/ajax";
import { CheckKey } from "src/rpc/api";

export const useCheckKey = (key: string): {keyValid: boolean, loading: boolean} => {
  const [loading, setLoading] = useState(true);
  const [keyValid, setKeyValid] = useState(false);

  const checkKey = async () => {
    try {
      await sendRequest(CheckKey, { key });
      setKeyValid(true);
    } catch(e) {
    }
    setLoading(false);
  }

  useEffect(() => {
    checkKey()
  }, []);

  return { keyValid, loading }
}