import { apiRequest } from "../../../lib/api";
import type { LoginRequest, LoginResponse } from "../../../types/auth";

export async function login(request: LoginRequest) {
  const response = await apiRequest<LoginResponse>("/auth/login", {
    method: "POST",
    body: request,
  });

  if (!response.data) {
    throw new Error("Login response did not include session data");
  }

  return response.data;
}
