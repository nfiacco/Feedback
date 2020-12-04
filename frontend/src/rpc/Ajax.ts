import { IEndpoint } from 'rpc/Api';

const IS_PROD = process.env.NODE_ENV === "production";
const ROOT_DOMAIN = IS_PROD ? "https://api.anonymousfeedback.app" : "http://localhost:8080";

export async function sendRequest<RequestType extends object | undefined, ResponseType>(
    endpoint: IEndpoint<RequestType, ResponseType>,
    payload?: RequestType,
): Promise<ResponseType | undefined | Error> {
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
        return new Error(response.statusText);
    }

    // not all AJAX requests have a response. the ones that do will be formatted as JSON
    const textBody = await response.text();
    if (textBody.length === 0) {
        return;
    }

    return JSON.parse(textBody);
}
