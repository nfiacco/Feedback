// eslint-disable-next-line @typescript-eslint/no-unused-vars
export interface IEndpoint<RequestType, ResponseType> {
    method: "GET" | "POST" | "DELETE" | "PUT";
    path: string;
}

export const Login: IEndpoint<LoginRequest, LoginResponse> = {
    method: "POST",
    path: "/login",
};

export const CheckSession: IEndpoint<undefined, CheckSessionResponse> = {
    method: "GET",
    path: "/check_session",
};

export const CheckKey: IEndpoint<{key: string}, undefined> = {
    method: "GET",
    path: "/check_key/:key",
};

export const SendFeedback: IEndpoint<SendRequest, undefined> = {
    method: "POST",
    path: "/send",
};

export interface SendRequest {
    feedback_key: string;
    escaped_content: string;
}

export interface LoginRequest {
    id_token: string;
}

export interface LoginResponse {
    first_name: string;
    last_name: string;
    feedback_key: string;
}

export interface CheckSessionResponse {
    first_name: string;
    last_name: string;
    feedback_key: string;
}