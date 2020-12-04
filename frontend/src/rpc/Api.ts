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

export interface LoginRequest {
    idtoken: string;
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