export type AuthUser = {
  id: string;
  name: string;
  email: string;
  role: string;
};

export type LoginRequest = {
  email: string;
  password: string;
};

export type LoginResponse = {
  access_token: string;
  token_type: "Bearer";
  expires_in: number;
  user: AuthUser;
};

export type AuthSession = {
  accessToken: string;
  tokenType: "Bearer";
  expiresIn: number;
  user: AuthUser;
};
