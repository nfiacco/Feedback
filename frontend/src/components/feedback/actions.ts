import { useEffect, useState } from "react";
import { sendRequest } from "src/rpc/ajax";
import { CheckKey, SendFeedback } from "src/rpc/api";
import { debounce } from "src/utils/debounce";

export const DEBOUNCE_DELAY = 200;

export type SendStatus = "SENT" | "SENDING" | "ERROR" | "NONE";

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

export function useSendFeedback(key: string, content: string) {
  const [sendStatus, setSendStatus] = useState("");

  // debounce the actual API call so we don't send multiple emails
  const toBeDebounced = async () => {
    try {
      const payload = {"feedback_key": key, "escaped_content": content};
      await sendRequest(SendFeedback, payload);
      setSendStatus("SENT");
    } catch (e) {
      setSendStatus("ERROR");
    }
  }

  const sendFeedback = async () => {
    setSendStatus("SENDING");
    debounce(toBeDebounced, DEBOUNCE_DELAY)();
  };

  const clearSendStatus = () => setSendStatus("NONE");
  return { sendFeedback, sendStatus, clearSendStatus };
}