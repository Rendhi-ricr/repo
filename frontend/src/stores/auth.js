/**
 * Auth Store
 * ===========
 * Svelte store untuk mengelola state autentikasi.
 * Menggunakan writable store dari Svelte.
 */

import { writable, derived } from "svelte/store";
import { APP_CONFIG } from "../config";

// Fungsi helper untuk membaca dari localStorage
function getStoredUser() {
    if (typeof localStorage === "undefined") return null;
    const user = localStorage.getItem(APP_CONFIG.USER_KEY);
    return user ? JSON.parse(user) : null;
}

function getStoredToken() {
    if (typeof localStorage === "undefined") return null;
    return localStorage.getItem(APP_CONFIG.TOKEN_KEY);
}

// State
export const currentUser = writable(getStoredUser());
export const authToken = writable(getStoredToken());
export const isLoading = writable(false);

// Derived stores (computed values)
export const isAuthenticated = derived(authToken, ($token) => !!$token);
export const isAdmin = derived(
    currentUser,
    ($user) => $user?.role === "admin"
);

// Actions
export function setAuth(token, user) {
    localStorage.setItem(APP_CONFIG.TOKEN_KEY, token);
    localStorage.setItem(APP_CONFIG.USER_KEY, JSON.stringify(user));
    authToken.set(token);
    currentUser.set(user);
}

export function clearAuth() {
    localStorage.removeItem(APP_CONFIG.TOKEN_KEY);
    localStorage.removeItem(APP_CONFIG.USER_KEY);
    authToken.set(null);
    currentUser.set(null);
}

export function setLoading(value) {
    isLoading.set(value);
}
