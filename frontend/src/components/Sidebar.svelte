<script>
    import { link } from "svelte-spa-router";
</script>

<aside
    class="w-72 bg-white dark:bg-slate-900 border-r border-slate-200 dark:border-slate-800 flex flex-col h-full shrink-0"
>
    <div class="p-6 flex flex-col gap-6">
        <!-- Brand -->
        <div class="flex items-center gap-3">
            <div
                class="bg-primary size-10 rounded-lg flex items-center justify-center text-white shadow-lg shadow-primary/20"
            >
                <span class="material-symbols-outlined">folder_managed</span>
            </div>
            <div class="flex flex-col">
                <h1
                    class="text-slate-900 dark:text-white text-base font-bold leading-tight"
                >
                    Admin Panel
                </h1>
                <p
                    class="text-slate-500 dark:text-slate-400 text-xs font-normal"
                >
                    Document System
                </p>
            </div>
        </div>

        <!-- Navigation -->
        <nav class="flex flex-col gap-1">
            <a
                href="#/"
                use:link
                class="flex items-center gap-3 px-3 py-2.5 rounded-lg text-slate-600 dark:text-slate-400 hover:bg-slate-100 dark:hover:bg-slate-800 transition-colors"
            >
                <span class="material-symbols-outlined text-xl">dashboard</span>
                <span class="text-sm font-medium">Dashboard</span>
            </a>
            <a
                href="#/documents"
                use:link
                class="flex items-center gap-3 px-3 py-2.5 rounded-lg bg-primary/10 text-primary dark:bg-primary/20"
            >
                <span
                    class="material-symbols-outlined text-xl"
                    style="font-variation-settings: 'FILL' 1">description</span
                >
                <span class="text-sm font-semibold">Kelola Dokumen</span>
            </a>
            <a
                href="#/users"
                use:link
                class="flex items-center gap-3 px-3 py-2.5 rounded-lg text-slate-600 dark:text-slate-400 hover:bg-slate-100 dark:hover:bg-slate-800 transition-colors"
            >
                <span class="material-symbols-outlined text-xl">group</span>
                <span class="text-sm font-medium">Manajemen User</span>
            </a>
            <a
                href="#/reports"
                use:link
                class="flex items-center gap-3 px-3 py-2.5 rounded-lg text-slate-600 dark:text-slate-400 hover:bg-slate-100 dark:hover:bg-slate-800 transition-colors"
            >
                <span class="material-symbols-outlined text-xl">assessment</span
                >
                <span class="text-sm font-medium">Laporan</span>
            </a>
        </nav>

        <hr class="border-slate-200 dark:border-slate-800" />

        <nav class="flex flex-col gap-1">
            <a
                href="#/settings"
                use:link
                class="flex items-center gap-3 px-3 py-2.5 rounded-lg text-slate-600 dark:text-slate-400 hover:bg-slate-100 dark:hover:bg-slate-800 transition-colors"
            >
                <span class="material-symbols-outlined text-xl">settings</span>
                <span class="text-sm font-medium">Pengaturan</span>
            </a>
            <a
                href="#/help"
                use:link
                class="flex items-center gap-3 px-3 py-2.5 rounded-lg text-slate-600 dark:text-slate-400 hover:bg-slate-100 dark:hover:bg-slate-800 transition-colors"
            >
                <span class="material-symbols-outlined text-xl">help</span>
                <span class="text-sm font-medium">Bantuan</span>
            </a>
        </nav>
    </div>

    <div class="mt-auto p-6 border-t border-slate-200 dark:border-slate-800">
        <div class="flex items-center gap-3 px-2">
            <div
                class="size-9 rounded-full bg-slate-200 dark:bg-slate-700 overflow-hidden shrink-0 border border-slate-200 dark:border-slate-700"
            >
                <img
                    class="w-full h-full object-cover"
                    src="https://ui-avatars.com/api/?name=Admin+Utama&background=135bec&color=fff"
                    alt="Admin"
                />
            </div>
            <div class="flex flex-col min-w-0">
                <span class="text-sm font-bold truncate">Admin Utama</span>
                <span class="text-xs text-slate-500 truncate"
                    >admin@company.com</span
                >
            </div>
            <button
                class="ml-auto text-slate-400 hover:text-red-500 transition-colors"
            >
                <span class="material-symbols-outlined">logout</span>
            </button>
        </div>
    </div>
</aside>
