export interface IEndpoint<RequestType extends object | undefined, ResponseType> {
    method: "GET" | "POST" | "DELETE" | "PUT";
    path: string;
}

export const Login: IEndpoint<LoginRequest, LoginResponse> = {
    method: "POST",
    path: "/login",
};

export const CheckSession: IEndpoint<undefined, undefined> = {
    method: "GET",
    path: "/check_session",
};

export interface LoginRequest {
    idtoken: string;
}

export interface LoginResponse {
    first_name: string;
    last_name: string;
}