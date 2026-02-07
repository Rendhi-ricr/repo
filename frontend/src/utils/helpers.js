/**
 * Utility Functions
 * ==================
 * Fungsi-fungsi helper yang digunakan di berbagai tempat.
 */

/**
 * Format tanggal ke format Indonesia
 * @param {Date|string} date - Tanggal yang akan diformat
 * @returns {string} - Tanggal dalam format Indonesia
 */
export function formatDate(date) {
    if (!date) return "-";
    const d = new Date(date);
    return d.toLocaleDateString("id-ID", {
        day: "2-digit",
        month: "long",
        year: "numeric",
    });
}

/**
 * Format tanggal dengan waktu
 * @param {Date|string} date - Tanggal yang akan diformat
 * @returns {string} - Tanggal dan waktu dalam format Indonesia
 */
export function formatDateTime(date) {
    if (!date) return "-";
    const d = new Date(date);
    return d.toLocaleDateString("id-ID", {
        day: "2-digit",
        month: "long",
        year: "numeric",
        hour: "2-digit",
        minute: "2-digit",
    });
}

/**
 * Format ukuran file ke format yang mudah dibaca
 * @param {number} bytes - Ukuran file dalam bytes
 * @returns {string} - Ukuran file yang sudah diformat
 */
export function formatFileSize(bytes) {
    if (!bytes) return "0 B";
    const sizes = ["B", "KB", "MB", "GB"];
    const i = Math.floor(Math.log(bytes) / Math.log(1024));
    return `${(bytes / Math.pow(1024, i)).toFixed(2)} ${sizes[i]}`;
}

/**
 * Truncate text dengan ellipsis
 * @param {string} text - Text yang akan dipotong
 * @param {number} maxLength - Panjang maksimal
 * @returns {string} - Text yang sudah dipotong
 */
export function truncate(text, maxLength = 50) {
    if (!text) return "";
    if (text.length <= maxLength) return text;
    return text.substring(0, maxLength) + "...";
}

/**
 * Debounce function untuk menunda eksekusi
 * @param {Function} func - Fungsi yang akan di-debounce
 * @param {number} wait - Waktu tunggu dalam ms
 * @returns {Function} - Fungsi yang sudah di-debounce
 */
export function debounce(func, wait = 300) {
    let timeout;
    return function executedFunction(...args) {
        const later = () => {
            clearTimeout(timeout);
            func(...args);
        };
        clearTimeout(timeout);
        timeout = setTimeout(later, wait);
    };
}

/**
 * Generate random string
 * @param {number} length - Panjang string
 * @returns {string} - Random string
 */
export function randomString(length = 8) {
    const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";
    let result = "";
    for (let i = 0; i < length; i++) {
        result += chars.charAt(Math.floor(Math.random() * chars.length));
    }
    return result;
}

/**
 * Copy text ke clipboard
 * @param {string} text - Text yang akan dicopy
 * @returns {Promise<boolean>} - Berhasil atau tidak
 */
export async function copyToClipboard(text) {
    try {
        await navigator.clipboard.writeText(text);
        return true;
    } catch (err) {
        console.error("Failed to copy:", err);
        return false;
    }
}

/**
 * Validasi email
 * @param {string} email - Email yang akan divalidasi
 * @returns {boolean} - Valid atau tidak
 */
export function isValidEmail(email) {
    const regex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
    return regex.test(email);
}

/**
 * Capitalize first letter
 * @param {string} str - String yang akan di-capitalize
 * @returns {string} - String dengan huruf pertama kapital
 */
export function capitalize(str) {
    if (!str) return "";
    return str.charAt(0).toUpperCase() + str.slice(1);
}
