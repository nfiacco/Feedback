import { IEndpoint } from 'rpc/Api';

const IS_PROD = process.env.NODE_ENV === "production";
const ROOT_DOMAIN = IS_PROD ? "https://api.anonymousfeedback.app" : "http://localhost:8080";

export async function sendRequest<RequestType, ResponseType>(
    endpoint: IEndpoint<RequestType, ResponseType>,
    payload?: RequestType,
): Promise<ResponseType> {
    const url = ROOT_DOMAIN + endpoint.path;
    let options: RequestInit = {
        method: endpoint.method,
        headers: new Headers({'Content-Type': 'application/json'}),
        credentials: 'include',
    };
    if (payload) {
        options.body = JSON.stringify(payload);
    }

    const response = await fetch(url, options);

    if (!response.ok) {
        throw new Error(response.statusText);
    }

    // not all AJAX requests have a response. the ones that do will be formatted as JSON
    return response.json().catch(() => {});
}
