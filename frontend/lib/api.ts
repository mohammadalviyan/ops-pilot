import type { ApiResponse } from "../types/api";

const defaultBaseURL = "http://localhost:8080/api/v1";

export class ApiError extends Error {
  status: number;
  errors?: ApiResponse<unknown>["errors"];

  constructor(message: string, status: number, errors?: ApiResponse<unknown>["errors"]) {
    super(message);
    this.name = "ApiError";
    this.status = status;
    this.errors = errors;
  }
}

type ApiRequestOptions = {
  method?: "GET" | "POST" | "PUT" | "PATCH" | "DELETE";
  body?: unknown;
  token?: string;
};

export async function apiRequest<T>(
  path: string,
  options: ApiRequestOptions = {},
): Promise<ApiResponse<T>> {
  const baseURL = process.env.NEXT_PUBLIC_API_BASE_URL || defaultBaseURL;
  const headers = new Headers({
    Accept: "application/json",
  });

  if (options.body !== undefined) {
    headers.set("Content-Type", "application/json");
  }

  if (options.token) {
    headers.set("Authorization", `Bearer ${options.token}`);
  }

  const response = await fetch(`${baseURL}${path}`, {
    method: options.method || "GET",
    headers,
    body: options.body === undefined ? undefined : JSON.stringify(options.body),
  });

  const payload = (await response.json().catch(() => ({
    success: false,
    message: "Unable to read server response",
  }))) as ApiResponse<T>;

  if (!response.ok || !payload.success) {
    throw new ApiError(
      payload.message || "Request failed",
      response.status,
      payload.errors,
    );
  }

  return payload;
}
