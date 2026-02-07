<script>
    import { onMount } from "svelte";
    import authService from "../../services/authService";

    let user = null;
    let activeTab = "profile";
    let saving = false;
    let message = { type: "", text: "" };

    // Profile form
    let profileForm = {
        name: "",
        email: "",
    };

    // Password form
    let passwordForm = {
        currentPassword: "",
        newPassword: "",
        confirmPassword: "",
    };

    // App settings
    let appSettings = {
        darkMode: false,
        notifications: true,
        emailNotifications: true,
        language: "id",
        timezone: "Asia/Jakarta",
    };

    onMount(() => {
        user = authService.getUser();
        if (user) {
            profileForm.name = user.name || "";
            profileForm.email = user.email || "";
        }

        // Load dark mode preference
        appSettings.darkMode =
            document.documentElement.classList.contains("dark");
    });

    function showMessage(type, text) {
        message = { type, text };
        setTimeout(() => {
            message = { type: "", text: "" };
        }, 3000);
    }

    async function saveProfile() {
        if (!profileForm.name || !profileForm.email) {
            showMessage("error", "Nama dan email harus diisi");
            return;
        }

        saving = true;
        try {
            // Simulate API call
            await new Promise((resolve) => setTimeout(resolve, 1000));

            // Update local storage
            const updatedUser = {
                ...user,
                name: profileForm.name,
                email: profileForm.email,
            };
            localStorage.setItem("auth_user", JSON.stringify(updatedUser));
            user = updatedUser;

            showMessage("success", "Profil berhasil diperbarui!");
        } catch (e) {
            showMessage("error", "Gagal menyimpan profil");
        } finally {
            saving = false;
        }
    }

    async function changePassword() {
        if (!passwordForm.currentPassword || !passwordForm.newPassword) {
            showMessage("error", "Semua field password harus diisi");
            return;
        }

        if (passwordForm.newPassword.length < 6) {
            showMessage("error", "Password baru minimal 6 karakter");
            return;
        }

        if (passwordForm.newPassword !== passwordForm.confirmPassword) {
            showMessage("error", "Konfirmasi password tidak cocok");
            return;
        }

        saving = true;
        try {
            // Simulate API call
            await new Promise((resolve) => setTimeout(resolve, 1000));

            passwordForm = {
                currentPassword: "",
                newPassword: "",
                confirmPassword: "",
            };
            showMessage("success", "Password berhasil diubah!");
        } catch (e) {
            showMessage("error", "Gagal mengubah password");
        } finally {
            saving = false;
        }
    }

    function toggleDarkMode() {
        appSettings.darkMode = !appSettings.darkMode;
        document.documentElement.classList.toggle("dark", appSettings.darkMode);
        localStorage.setItem("darkMode", String(appSettings.darkMode));
    }

    function saveAppSettings() {
        showMessage("success", "Pengaturan berhasil disimpan!");
    }

    const tabs = [
        { id: "profile", label: "Profil", icon: "person" },
        { id: "security", label: "Keamanan", icon: "lock" },
        { id: "appearance", label: "Tampilan", icon: "palette" },
        { id: "notifications", label: "Notifikasi", icon: "notifications" },
        { id: "system", label: "Sistem", icon: "settings" },
    ];
</script>

<div class="max-w-5xl mx-auto flex flex-col gap-6">
    <!-- Header -->
    <div class="flex flex-col gap-1">
        <h2
            class="text-3xl font-black text-slate-900 dark:text-white tracking-tight"
        >
            Pengaturan
        </h2>
        <p class="text-slate-500 dark:text-slate-400">
            Kelola preferensi akun dan pengaturan aplikasi.
        </p>
    </div>

    <!-- Message Toast -->
    {#if message.text}
        <div
            class="fixed top-4 right-4 z-50 px-6 py-3 rounded-lg shadow-lg flex items-center gap-3 animate-slide-in
                {message.type === 'success'
                ? 'bg-emerald-500 text-white'
                : 'bg-red-500 text-white'}"
        >
            <span class="material-symbols-outlined">
                {message.type === "success" ? "check_circle" : "error"}
            </span>
            {message.text}
        </div>
    {/if}

    <div class="grid grid-cols-1 lg:grid-cols-4 gap-6">
        <!-- Sidebar Tabs -->
        <div class="lg:col-span-1">
            <nav
                class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 p-2 shadow-sm"
            >
                {#each tabs as tab}
                    <button
                        on:click={() => (activeTab = tab.id)}
                        class="w-full flex items-center gap-3 px-4 py-3 rounded-xl text-left transition-all
                            {activeTab === tab.id
                            ? 'bg-primary text-white shadow-lg shadow-primary/25'
                            : 'text-slate-600 dark:text-slate-400 hover:bg-slate-100 dark:hover:bg-slate-800'}"
                    >
                        <span class="material-symbols-outlined text-xl"
                            >{tab.icon}</span
                        >
                        <span class="font-medium">{tab.label}</span>
                    </button>
                {/each}
            </nav>
        </div>

        <!-- Content Area -->
        <div class="lg:col-span-3">
            <div
                class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 shadow-sm overflow-hidden"
            >
                <!-- Profile Tab -->
                {#if activeTab === "profile"}
                    <div
                        class="p-6 border-b border-slate-200 dark:border-slate-800"
                    >
                        <h3
                            class="text-lg font-bold text-slate-900 dark:text-white"
                        >
                            Informasi Profil
                        </h3>
                        <p class="text-sm text-slate-500">
                            Perbarui informasi akun Anda
                        </p>
                    </div>
                    <div class="p-6 space-y-6">
                        <!-- Avatar -->
                        <div class="flex items-center gap-6">
                            <div
                                class="w-20 h-20 rounded-full bg-gradient-to-br from-primary to-blue-600 flex items-center justify-center text-white text-3xl font-bold shadow-lg"
                            >
                                {user?.name?.charAt(0).toUpperCase() || "U"}
                            </div>
                            <div>
                                <button
                                    class="px-4 py-2 bg-primary text-white font-medium rounded-lg hover:bg-primary/90 transition-colors"
                                >
                                    Upload Foto
                                </button>
                                <p class="text-sm text-slate-500 mt-2">
                                    JPG, PNG maksimal 2MB
                                </p>
                            </div>
                        </div>

                        <!-- Form -->
                        <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                            <div>
                                <label
                                    class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2"
                                >
                                    Nama Lengkap
                                </label>
                                <input
                                    type="text"
                                    bind:value={profileForm.name}
                                    class="w-full px-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl focus:ring-2 focus:ring-primary focus:border-transparent transition-all"
                                    placeholder="Masukkan nama"
                                />
                            </div>
                            <div>
                                <label
                                    class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2"
                                >
                                    Email
                                </label>
                                <input
                                    type="email"
                                    bind:value={profileForm.email}
                                    class="w-full px-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl focus:ring-2 focus:ring-primary focus:border-transparent transition-all"
                                    placeholder="email@example.com"
                                />
                            </div>
                        </div>

                        <div class="flex justify-end pt-4">
                            <button
                                on:click={saveProfile}
                                disabled={saving}
                                class="flex items-center gap-2 px-6 py-3 bg-primary text-white font-bold rounded-xl shadow-lg shadow-primary/25 hover:bg-primary/90 transition-all disabled:opacity-50"
                            >
                                {#if saving}
                                    <div
                                        class="w-5 h-5 border-2 border-white/30 border-t-white rounded-full animate-spin"
                                    ></div>
                                {:else}
                                    <span class="material-symbols-outlined"
                                        >save</span
                                    >
                                {/if}
                                Simpan Perubahan
                            </button>
                        </div>
                    </div>
                {/if}

                <!-- Security Tab -->
                {#if activeTab === "security"}
                    <div
                        class="p-6 border-b border-slate-200 dark:border-slate-800"
                    >
                        <h3
                            class="text-lg font-bold text-slate-900 dark:text-white"
                        >
                            Keamanan Akun
                        </h3>
                        <p class="text-sm text-slate-500">
                            Kelola password dan keamanan akun
                        </p>
                    </div>
                    <div class="p-6 space-y-6">
                        <div class="space-y-4">
                            <div>
                                <label
                                    class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2"
                                >
                                    Password Saat Ini
                                </label>
                                <input
                                    type="password"
                                    bind:value={passwordForm.currentPassword}
                                    class="w-full px-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl focus:ring-2 focus:ring-primary focus:border-transparent transition-all"
                                    placeholder="••••••••"
                                />
                            </div>
                            <div>
                                <label
                                    class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2"
                                >
                                    Password Baru
                                </label>
                                <input
                                    type="password"
                                    bind:value={passwordForm.newPassword}
                                    class="w-full px-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl focus:ring-2 focus:ring-primary focus:border-transparent transition-all"
                                    placeholder="Minimal 6 karakter"
                                />
                            </div>
                            <div>
                                <label
                                    class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2"
                                >
                                    Konfirmasi Password Baru
                                </label>
                                <input
                                    type="password"
                                    bind:value={passwordForm.confirmPassword}
                                    class="w-full px-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl focus:ring-2 focus:ring-primary focus:border-transparent transition-all"
                                    placeholder="Ulangi password baru"
                                />
                            </div>
                        </div>

                        <div class="flex justify-end pt-4">
                            <button
                                on:click={changePassword}
                                disabled={saving}
                                class="flex items-center gap-2 px-6 py-3 bg-primary text-white font-bold rounded-xl shadow-lg shadow-primary/25 hover:bg-primary/90 transition-all disabled:opacity-50"
                            >
                                {#if saving}
                                    <div
                                        class="w-5 h-5 border-2 border-white/30 border-t-white rounded-full animate-spin"
                                    ></div>
                                {:else}
                                    <span class="material-symbols-outlined"
                                        >lock</span
                                    >
                                {/if}
                                Ubah Password
                            </button>
                        </div>

                        <hr class="border-slate-200 dark:border-slate-800" />

                        <!-- Two Factor Auth -->
                        <div
                            class="flex items-center justify-between p-4 bg-slate-50 dark:bg-slate-800 rounded-xl"
                        >
                            <div class="flex items-center gap-4">
                                <div
                                    class="p-3 bg-amber-100 dark:bg-amber-900/30 rounded-lg"
                                >
                                    <span
                                        class="material-symbols-outlined text-amber-600"
                                        >security</span
                                    >
                                </div>
                                <div>
                                    <h4
                                        class="font-semibold text-slate-900 dark:text-white"
                                    >
                                        Autentikasi Dua Faktor
                                    </h4>
                                    <p class="text-sm text-slate-500">
                                        Tambahkan lapisan keamanan ekstra
                                    </p>
                                </div>
                            </div>
                            <button
                                class="px-4 py-2 bg-white dark:bg-slate-700 border border-slate-200 dark:border-slate-600 rounded-lg font-medium text-slate-700 dark:text-slate-300 hover:bg-slate-100 dark:hover:bg-slate-600 transition-colors"
                            >
                                Aktifkan
                            </button>
                        </div>
                    </div>
                {/if}

                <!-- Appearance Tab -->
                {#if activeTab === "appearance"}
                    <div
                        class="p-6 border-b border-slate-200 dark:border-slate-800"
                    >
                        <h3
                            class="text-lg font-bold text-slate-900 dark:text-white"
                        >
                            Tampilan
                        </h3>
                        <p class="text-sm text-slate-500">
                            Sesuaikan tampilan aplikasi
                        </p>
                    </div>
                    <div class="p-6 space-y-6">
                        <!-- Dark Mode Toggle -->
                        <div
                            class="flex items-center justify-between p-4 bg-slate-50 dark:bg-slate-800 rounded-xl"
                        >
                            <div class="flex items-center gap-4">
                                <div
                                    class="p-3 bg-slate-200 dark:bg-slate-700 rounded-lg"
                                >
                                    <span
                                        class="material-symbols-outlined text-slate-600 dark:text-slate-300"
                                    >
                                        {appSettings.darkMode
                                            ? "dark_mode"
                                            : "light_mode"}
                                    </span>
                                </div>
                                <div>
                                    <h4
                                        class="font-semibold text-slate-900 dark:text-white"
                                    >
                                        Mode Gelap
                                    </h4>
                                    <p class="text-sm text-slate-500">
                                        {appSettings.darkMode
                                            ? "Aktif"
                                            : "Nonaktif"}
                                    </p>
                                </div>
                            </div>
                            <button
                                on:click={toggleDarkMode}
                                class="relative w-14 h-7 rounded-full transition-colors {appSettings.darkMode
                                    ? 'bg-primary'
                                    : 'bg-slate-300'}"
                            >
                                <div
                                    class="absolute top-1 w-5 h-5 bg-white rounded-full shadow transition-transform {appSettings.darkMode
                                        ? 'translate-x-8'
                                        : 'translate-x-1'}"
                                ></div>
                            </button>
                        </div>

                        <!-- Language -->
                        <div>
                            <label
                                class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2"
                            >
                                Bahasa
                            </label>
                            <select
                                bind:value={appSettings.language}
                                class="w-full px-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl focus:ring-2 focus:ring-primary focus:border-transparent transition-all"
                            >
                                <option value="id">Bahasa Indonesia</option>
                                <option value="en">English</option>
                            </select>
                        </div>

                        <!-- Timezone -->
                        <div>
                            <label
                                class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2"
                            >
                                Zona Waktu
                            </label>
                            <select
                                bind:value={appSettings.timezone}
                                class="w-full px-4 py-3 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 rounded-xl focus:ring-2 focus:ring-primary focus:border-transparent transition-all"
                            >
                                <option value="Asia/Jakarta"
                                    >WIB (Jakarta)</option
                                >
                                <option value="Asia/Makassar"
                                    >WITA (Makassar)</option
                                >
                                <option value="Asia/Jayapura"
                                    >WIT (Jayapura)</option
                                >
                            </select>
                        </div>
                    </div>
                {/if}

                <!-- Notifications Tab -->
                {#if activeTab === "notifications"}
                    <div
                        class="p-6 border-b border-slate-200 dark:border-slate-800"
                    >
                        <h3
                            class="text-lg font-bold text-slate-900 dark:text-white"
                        >
                            Notifikasi
                        </h3>
                        <p class="text-sm text-slate-500">
                            Kelola preferensi notifikasi
                        </p>
                    </div>
                    <div class="p-6 space-y-4">
                        <div
                            class="flex items-center justify-between p-4 bg-slate-50 dark:bg-slate-800 rounded-xl"
                        >
                            <div>
                                <h4
                                    class="font-semibold text-slate-900 dark:text-white"
                                >
                                    Notifikasi Push
                                </h4>
                                <p class="text-sm text-slate-500">
                                    Terima notifikasi di browser
                                </p>
                            </div>
                            <button
                                on:click={() =>
                                    (appSettings.notifications =
                                        !appSettings.notifications)}
                                class="relative w-14 h-7 rounded-full transition-colors {appSettings.notifications
                                    ? 'bg-primary'
                                    : 'bg-slate-300'}"
                            >
                                <div
                                    class="absolute top-1 w-5 h-5 bg-white rounded-full shadow transition-transform {appSettings.notifications
                                        ? 'translate-x-8'
                                        : 'translate-x-1'}"
                                ></div>
                            </button>
                        </div>

                        <div
                            class="flex items-center justify-between p-4 bg-slate-50 dark:bg-slate-800 rounded-xl"
                        >
                            <div>
                                <h4
                                    class="font-semibold text-slate-900 dark:text-white"
                                >
                                    Notifikasi Email
                                </h4>
                                <p class="text-sm text-slate-500">
                                    Terima update via email
                                </p>
                            </div>
                            <button
                                on:click={() =>
                                    (appSettings.emailNotifications =
                                        !appSettings.emailNotifications)}
                                class="relative w-14 h-7 rounded-full transition-colors {appSettings.emailNotifications
                                    ? 'bg-primary'
                                    : 'bg-slate-300'}"
                            >
                                <div
                                    class="absolute top-1 w-5 h-5 bg-white rounded-full shadow transition-transform {appSettings.emailNotifications
                                        ? 'translate-x-8'
                                        : 'translate-x-1'}"
                                ></div>
                            </button>
                        </div>
                    </div>
                {/if}

                <!-- System Tab -->
                {#if activeTab === "system"}
                    <div
                        class="p-6 border-b border-slate-200 dark:border-slate-800"
                    >
                        <h3
                            class="text-lg font-bold text-slate-900 dark:text-white"
                        >
                            Informasi Sistem
                        </h3>
                        <p class="text-sm text-slate-500">
                            Detail teknis aplikasi
                        </p>
                    </div>
                    <div class="p-6 space-y-4">
                        <div class="grid grid-cols-2 gap-4">
                            <div
                                class="p-4 bg-slate-50 dark:bg-slate-800 rounded-xl"
                            >
                                <span class="text-sm text-slate-500"
                                    >Versi Aplikasi</span
                                >
                                <p
                                    class="text-lg font-bold text-slate-900 dark:text-white"
                                >
                                    1.0.0
                                </p>
                            </div>
                            <div
                                class="p-4 bg-slate-50 dark:bg-slate-800 rounded-xl"
                            >
                                <span class="text-sm text-slate-500"
                                    >Status Server</span
                                >
                                <p
                                    class="text-lg font-bold text-emerald-500 flex items-center gap-2"
                                >
                                    <span
                                        class="w-2 h-2 bg-emerald-500 rounded-full animate-pulse"
                                    ></span>
                                    Online
                                </p>
                            </div>
                            <div
                                class="p-4 bg-slate-50 dark:bg-slate-800 rounded-xl"
                            >
                                <span class="text-sm text-slate-500"
                                    >Database</span
                                >
                                <p
                                    class="text-lg font-bold text-slate-900 dark:text-white"
                                >
                                    PostgreSQL
                                </p>
                            </div>
                            <div
                                class="p-4 bg-slate-50 dark:bg-slate-800 rounded-xl"
                            >
                                <span class="text-sm text-slate-500"
                                    >Backend</span
                                >
                                <p
                                    class="text-lg font-bold text-slate-900 dark:text-white"
                                >
                                    Go 1.21
                                </p>
                            </div>
                        </div>

                        <hr class="border-slate-200 dark:border-slate-800" />

                        <div
                            class="p-4 bg-red-50 dark:bg-red-900/20 rounded-xl border border-red-200 dark:border-red-800"
                        >
                            <h4
                                class="font-semibold text-red-700 dark:text-red-400 mb-2"
                            >
                                Zona Berbahaya
                            </h4>
                            <p
                                class="text-sm text-red-600 dark:text-red-400 mb-4"
                            >
                                Tindakan berikut tidak dapat dibatalkan. Harap
                                berhati-hati.
                            </p>
                            <div class="flex gap-3">
                                <button
                                    class="px-4 py-2 bg-white dark:bg-slate-800 border border-red-300 dark:border-red-700 text-red-600 dark:text-red-400 font-medium rounded-lg hover:bg-red-50 dark:hover:bg-red-900/30 transition-colors"
                                >
                                    Clear Cache
                                </button>
                                <button
                                    class="px-4 py-2 bg-red-600 text-white font-medium rounded-lg hover:bg-red-700 transition-colors"
                                >
                                    Reset Aplikasi
                                </button>
                            </div>
                        </div>
                    </div>
                {/if}
            </div>
        </div>
    </div>
</div>

<style>
    @keyframes slide-in {
        from {
            transform: translateX(100%);
            opacity: 0;
        }
        to {
            transform: translateX(0);
            opacity: 1;
        }
    }
    .animate-slide-in {
        animation: slide-in 0.3s ease-out;
    }
</style>
