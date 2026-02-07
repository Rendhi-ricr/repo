/**
 * Konfigurasi aplikasi
 * =====================
 * File ini berisi semua konfigurasi yang digunakan di aplikasi.
 * Ubah nilai di sini untuk menyesuaikan dengan environment.
 */

// URL API Backend
// Untuk development: http://localhost:8080
// Untuk production: ganti dengan URL server production
export const API_BASE_URL = "http://localhost:8080";

// API Endpoints
export const API_ENDPOINTS = {
    // Auth
    AUTH_LOGIN: `${API_BASE_URL}/api/auth/login`,
    AUTH_REGISTER: `${API_BASE_URL}/api/auth/register`,
    AUTH_ME: `${API_BASE_URL}/api/auth/me`,

    // Users
    USERS: `${API_BASE_URL}/api/users`,
    USER_BY_ID: (id) => `${API_BASE_URL}/api/users/${id}`,

    // Documents
    DOCUMENTS: `${API_BASE_URL}/api/documents`,
    DOCUMENT_BY_ID: (id) => `${API_BASE_URL}/api/documents/${id}`,
    DOCUMENT_PAGES: (id) => `${API_BASE_URL}/api/documents/pages/${id}`,
    DOCUMENT_DOWNLOAD: (id) => `${API_BASE_URL}/download/${id}`,
    DOCUMENT_PREVIEW: (id, page) => `${API_BASE_URL}/preview/split/${id}/${page}`,
};

// App Configuration
export const APP_CONFIG = {
    APP_NAME: "ScholarHub",
    APP_DESCRIPTION: "Repository Dokumen Akademik",

    // Storage keys
    TOKEN_KEY: "auth_token",
    USER_KEY: "auth_user",

    // Pagination
    DEFAULT_PAGE_SIZE: 10,

    // File upload
    MAX_FILE_SIZE: 10 * 1024 * 1024, // 10 MB
    ALLOWED_FILE_TYPES: [".pdf", ".doc", ".docx"],
};

// Route paths
export const ROUTES = {
    // Public
    HOME: "/",
    LANDING: "/landing",
    BROWSE: "/browse",
    ABOUT: "/about",

    // Auth
    LOGIN: "/login",
    REGISTER: "/register",

    // Admin
    ADMIN: "/admin",
    DASHBOARD: "/admin/dashboard",
    DOCUMENTS: "/admin/documents",
    DOCUMENTS_ADD: "/admin/documents/add",
    DOCUMENTS_EDIT: (id) => `/admin/documents/edit/${id}`,
    USERS: "/admin/users",
    REPORTS: "/reports",
    SETTINGS: "/settings",
};
