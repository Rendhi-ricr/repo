<script>
    import authService from "../../services/authService";

    let name = "";
    let email = "";
    let password = "";
    let confirmPassword = "";
    let loading = false;
    let error = "";

    async function handleSubmit() {
        if (!name || !email || !password || !confirmPassword) {
            error = "Please fill in all fields";
            return;
        }

        if (password !== confirmPassword) {
            error = "Passwords do not match";
            return;
        }

        if (password.length < 6) {
            error = "Password must be at least 6 characters";
            return;
        }

        loading = true;
        error = "";

        try {
            await authService.register(name, email, password);
            window.location.hash = "/admin";
            window.location.reload();
        } catch (err) {
            error = err.message;
        } finally {
            loading = false;
        }
    }
</script>

<section
    class="min-h-screen flex items-center justify-center px-4 py-12 bg-gradient-to-br from-background-dark via-surface-dark to-background-dark"
>
    <!-- Background decorations -->
    <div class="absolute inset-0 overflow-hidden">
        <div
            class="absolute top-0 right-1/4 w-96 h-96 bg-purple-500/10 rounded-full blur-3xl"
        ></div>
        <div
            class="absolute bottom-0 left-1/4 w-96 h-96 bg-primary/10 rounded-full blur-3xl"
        ></div>
    </div>

    <div class="relative w-full max-w-md">
        <!-- Card -->
        <div
            class="bg-white dark:bg-surface-highlight border border-slate-200 dark:border-slate-700 rounded-2xl shadow-2xl p-8"
        >
            <!-- Logo -->
            <div class="text-center mb-8">
                <a
                    href="#/"
                    class="inline-flex items-center gap-2 text-slate-900 dark:text-white"
                >
                    <span
                        class="material-symbols-outlined text-primary text-4xl"
                        >local_library</span
                    >
                    <span class="text-2xl font-bold">ScholarHub</span>
                </a>
                <h1
                    class="mt-6 text-2xl font-bold text-slate-900 dark:text-white"
                >
                    Create Account
                </h1>
                <p class="mt-2 text-slate-500 dark:text-slate-400">
                    Join our community of researchers
                </p>
            </div>

            <!-- Error Message -->
            {#if error}
                <div
                    class="mb-6 p-4 bg-red-100 dark:bg-red-900/30 border border-red-200 dark:border-red-800 rounded-lg"
                >
                    <p
                        class="text-red-600 dark:text-red-400 text-sm flex items-center gap-2"
                    >
                        <span class="material-symbols-outlined text-lg"
                            >error</span
                        >
                        {error}
                    </p>
                </div>
            {/if}

            <!-- Form -->
            <form on:submit|preventDefault={handleSubmit} class="space-y-4">
                <div>
                    <label
                        for="name"
                        class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2"
                    >
                        Full Name
                    </label>
                    <div class="relative">
                        <div
                            class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none"
                        >
                            <span
                                class="material-symbols-outlined text-slate-400 text-xl"
                                >person</span
                            >
                        </div>
                        <input
                            type="text"
                            id="name"
                            bind:value={name}
                            class="w-full h-12 pl-10 pr-4 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-600 rounded-lg text-slate-900 dark:text-white placeholder:text-slate-400 focus:outline-none focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all"
                            placeholder="Enter your full name"
                            required
                        />
                    </div>
                </div>

                <div>
                    <label
                        for="email"
                        class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2"
                    >
                        Email Address
                    </label>
                    <div class="relative">
                        <div
                            class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none"
                        >
                            <span
                                class="material-symbols-outlined text-slate-400 text-xl"
                                >mail</span
                            >
                        </div>
                        <input
                            type="email"
                            id="email"
                            bind:value={email}
                            class="w-full h-12 pl-10 pr-4 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-600 rounded-lg text-slate-900 dark:text-white placeholder:text-slate-400 focus:outline-none focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all"
                            placeholder="Enter your email"
                            required
                        />
                    </div>
                </div>

                <div>
                    <label
                        for="password"
                        class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2"
                    >
                        Password
                    </label>
                    <div class="relative">
                        <div
                            class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none"
                        >
                            <span
                                class="material-symbols-outlined text-slate-400 text-xl"
                                >lock</span
                            >
                        </div>
                        <input
                            type="password"
                            id="password"
                            bind:value={password}
                            class="w-full h-12 pl-10 pr-4 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-600 rounded-lg text-slate-900 dark:text-white placeholder:text-slate-400 focus:outline-none focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all"
                            placeholder="Create a password (min 6 characters)"
                            required
                        />
                    </div>
                </div>

                <div>
                    <label
                        for="confirmPassword"
                        class="block text-sm font-medium text-slate-700 dark:text-slate-300 mb-2"
                    >
                        Confirm Password
                    </label>
                    <div class="relative">
                        <div
                            class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none"
                        >
                            <span
                                class="material-symbols-outlined text-slate-400 text-xl"
                                >lock</span
                            >
                        </div>
                        <input
                            type="password"
                            id="confirmPassword"
                            bind:value={confirmPassword}
                            class="w-full h-12 pl-10 pr-4 bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-600 rounded-lg text-slate-900 dark:text-white placeholder:text-slate-400 focus:outline-none focus:ring-2 focus:ring-primary/50 focus:border-primary transition-all"
                            placeholder="Confirm your password"
                            required
                        />
                    </div>
                </div>

                <div class="pt-2">
                    <label class="flex items-start gap-3 cursor-pointer">
                        <input
                            type="checkbox"
                            required
                            class="w-4 h-4 mt-0.5 rounded border-slate-300 text-primary focus:ring-primary"
                        />
                        <span
                            class="text-sm text-slate-600 dark:text-slate-400"
                        >
                            I agree to the <a
                                href="#/terms"
                                class="text-primary hover:underline"
                                >Terms of Service</a
                            >
                            and
                            <a
                                href="#/privacy"
                                class="text-primary hover:underline"
                                >Privacy Policy</a
                            >
                        </span>
                    </label>
                </div>

                <button
                    type="submit"
                    disabled={loading}
                    class="w-full h-12 bg-primary hover:bg-primary/90 disabled:bg-primary/50 text-white font-bold rounded-lg transition-all flex items-center justify-center gap-2 shadow-lg shadow-primary/25 mt-6"
                >
                    {#if loading}
                        <svg class="animate-spin h-5 w-5" viewBox="0 0 24 24">
                            <circle
                                class="opacity-25"
                                cx="12"
                                cy="12"
                                r="10"
                                stroke="currentColor"
                                stroke-width="4"
                                fill="none"
                            ></circle>
                            <path
                                class="opacity-75"
                                fill="currentColor"
                                d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
                            ></path>
                        </svg>
                        <span>Creating account...</span>
                    {:else}
                        <span>Create Account</span>
                        <span class="material-symbols-outlined text-lg"
                            >arrow_forward</span
                        >
                    {/if}
                </button>
            </form>

            <!-- Divider -->
            <div class="my-6 flex items-center gap-4">
                <div class="flex-1 h-px bg-slate-200 dark:bg-slate-700"></div>
                <span class="text-sm text-slate-400">or</span>
                <div class="flex-1 h-px bg-slate-200 dark:bg-slate-700"></div>
            </div>

            <!-- Login Link -->
            <p class="text-center text-slate-600 dark:text-slate-400">
                Already have an account?
                <a
                    href="#/login"
                    class="text-primary font-medium hover:underline">Sign in</a
                >
            </p>
        </div>

        <!-- Back to home -->
        <p class="mt-6 text-center">
            <a
                href="#/"
                class="text-slate-500 dark:text-slate-400 hover:text-primary transition-colors inline-flex items-center gap-1"
            >
                <span class="material-symbols-outlined text-sm">arrow_back</span
                >
                Back to home
            </a>
        </p>
    </div>
</section>

<style>
    .bg-surface-dark {
        background-color: #192633;
    }
    .bg-surface-highlight {
        background-color: #233648;
    }
</style>
