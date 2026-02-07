/**
 * Stores Index
 * =============
 * Export semua stores dari satu tempat.
 */

// Auth store
export {
    currentUser,
    authToken,
    isAuthenticated,
    isAdmin,
    isLoading,
    setAuth,
    clearAuth,
    setLoading,
} from "./auth.js";
