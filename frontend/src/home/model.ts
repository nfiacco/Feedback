const INITIAL_LOGIN_STATE: LoginState = {
  authenticated: false,
};

export interface LoginState {
  authenticated: boolean;
  feedbackKey?: string;
}

export type LoginAction = 
| {
  type: "login.authenticated",
  feedbackKey: string,
};

export function loginReducer(state: LoginState = INITIAL_LOGIN_STATE, action: LoginAction): LoginState {
  switch (action.type) {
    case "login.authenticated":
    return {
      ...state,
      authenticated: true,
      feedbackKey: action.feedbackKey,
    }
    default:
    return state;
  }
}