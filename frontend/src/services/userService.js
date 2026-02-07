import authService from "./authService";

const API_BASE = "http://localhost:8080";

class UserService {
    async getUsers() {
        const response = await fetch(`${API_BASE}/api/users`, {
            headers: authService.getAuthHeaders(),
        });

        if (!response.ok) {
            const error = await response.json();
            throw new Error(error.error || "Failed to fetch users");
        }

        return response.json();
    }

    async getUserById(id) {
        const response = await fetch(`${API_BASE}/api/users/${id}`, {
            headers: authService.getAuthHeaders(),
        });

        if (!response.ok) {
            const error = await response.json();
            throw new Error(error.error || "Failed to fetch user");
        }

        return response.json();
    }

    async createUser(userData) {
        const response = await fetch(`${API_BASE}/api/users`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                ...authService.getAuthHeaders(),
            },
            body: JSON.stringify(userData),
        });

        if (!response.ok) {
            const error = await response.json();
            throw new Error(error.error || "Failed to create user");
        }

        return response.json();
    }

    async updateUser(id, userData) {
        const response = await fetch(`${API_BASE}/api/users/${id}`, {
            method: "PUT",
            headers: {
                "Content-Type": "application/json",
                ...authService.getAuthHeaders(),
            },
            body: JSON.stringify(userData),
        });

        if (!response.ok) {
            const error = await response.json();
            throw new Error(error.error || "Failed to update user");
        }

        return response.json();
    }

    async deleteUser(id) {
        const response = await fetch(`${API_BASE}/api/users/${id}`, {
            method: "DELETE",
            headers: authService.getAuthHeaders(),
        });

        if (!response.ok) {
            const error = await response.json();
            throw new Error(error.error || "Failed to delete user");
        }

        return response.json();
    }
}

export const userService = new UserService();
export default userService;
