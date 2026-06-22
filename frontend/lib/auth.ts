import type { AuthSession } from "../types/auth";

const storageKey = "opspilot.auth.session";

export function saveAuthSession(session: AuthSession) {
  if (typeof window === "undefined") {
    return;
  }

  window.localStorage.setItem(storageKey, JSON.stringify(session));
}

export function readAuthSession(): AuthSession | null {
  if (typeof window === "undefined") {
    return null;
  }

  const rawSession = window.localStorage.getItem(storageKey);
  if (!rawSession) {
    return null;
  }

  try {
    return JSON.parse(rawSession) as AuthSession;
  } catch {
    window.localStorage.removeItem(storageKey);
    return null;
  }
}

export function clearAuthSession() {
  if (typeof window === "undefined") {
    return;
  }

  window.localStorage.removeItem(storageKey);
}
