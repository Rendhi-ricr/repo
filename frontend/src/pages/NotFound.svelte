<script>
    import { link } from "svelte-spa-router";
</script>

<div
    class="flex flex-col items-center justify-center min-h-[60vh] text-center p-8"
>
    <div class="bg-slate-100 dark:bg-slate-800 p-6 rounded-full mb-6">
        <span class="material-symbols-outlined text-6xl text-slate-400"
            >sentiment_dissatisfied</span
        >
    </div>

    <h1 class="text-4xl font-black text-slate-900 dark:text-white mb-2">404</h1>
    <h2 class="text-xl font-bold text-slate-700 dark:text-slate-300 mb-4">
        Halaman Tidak Ditemukan
    </h2>

    <p class="text-slate-500 dark:text-slate-400 max-w-md mb-8">
        Maaf, halaman yang Anda cari tidak dapat ditemukan atau telah
        dipindahkan.
    </p>

    <a
        href="#/"
        use:link
        class="flex items-center gap-2 px-6 py-3 bg-primary text-white font-bold rounded-lg shadow-lg shadow-primary/25 hover:bg-primary/90 transition-all"
    >
        <span class="material-symbols-outlined">home</span>
        <span>Kembali ke Dashboard</span>
    </a>
</div>
