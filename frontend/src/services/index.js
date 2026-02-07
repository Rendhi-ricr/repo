/**
 * Services Index
 * ===============
 * Export semua services dari satu tempat.
 * 
 * Penggunaan:
 * ```javascript
 * import { authService, documentService, userService } from '@/services';
 * ```
 */

// Auth Service
export { authService } from "./authService.js";

// Document Service
export {
    documentService,
    getDocuments,
    getDocumentById,
    createDocument,
    updateDocument,
    deleteDocument,
    downloadDocument,
    getDocumentPages,
} from "./documentService.js";

// User Service
export {
    userService,
    getUsers,
    getUserById,
    createUser,
    updateUser,
    deleteUser,
} from "./userService.js";
