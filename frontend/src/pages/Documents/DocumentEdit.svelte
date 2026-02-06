<script>
    import { onMount } from "svelte";
    import { link } from "svelte-spa-router";
    import {
        getDocumentById,
        updateDocument,
    } from "../../services/documentService.js";

    export let params = {};

    let title = "";
    let author = "";
    let fileType = "";
    let status = "draft";
    let file = null;
    let currentFileName = "";
    let confirm = false;

    let loading = false;
    let loadingData = true;
    let error = "";

    onMount(async () => {
        await loadDocument();
    });

    async function loadDocument() {
        try {
            loadingData = true;
            const doc = await getDocumentById(params.id);
            title = doc.judul;
            author = doc.penulis;
            fileType = doc.jenis_file;
            status = doc.status;
            currentFileName = "File saat ini tersimpan di server";
        } catch (e) {
            error = "Gagal memuat data dokumen: " + e.message;
        } finally {
            loadingData = false;
        }
    }

    function handleFileChange(e) {
        file = e.target.files[0];
    }

    async function handleSubmit() {
        error = "";

        if (!title || !author || !fileType || !confirm) {
            error = "Semua field wajib diisi dan dikonfirmasi.";
            return;
        }

        loading = true;

        const formData = new FormData();
        formData.append("title", title);
        formData.append("author", author);
        formData.append("category", fileType);
        formData.append("status", status);

        if (file) {
            formData.append("file", file);
        }

        try {
            await updateDocument(params.id, formData);
            alert("Dokumen berhasil diupdate");
            window.location.href = "#/documents";
        } catch (e) {
            error = e.message;
        } finally {
            loading = false;
        }
    }
</script>

<div class="max-w-4xl mx-auto flex flex-col gap-6">
    <!-- Breadcrumb -->
    <nav class="flex items-center gap-2 text-sm">
        <a
            href="#/"
            use:link
            class="text-slate-500 hover:text-primary transition-colors"
            >Dashboard</a
        >
        <span class="text-slate-400">/</span>
        <a
            href="#/documents"
            use:link
            class="text-slate-500 hover:text-primary transition-colors"
            >Kelola Dokumen</a
        >
        <span class="text-slate-400">/</span>
        <span class="text-slate-900 dark:text-white font-medium"
            >Edit Dokumen</span
        >
    </nav>

    <!-- Header -->
    <div class="flex flex-col gap-1">
        <h2
            class="text-3xl font-black text-slate-900 dark:text-white tracking-tight"
        >
            Edit Dokumen
        </h2>
        <p class="text-slate-500 dark:text-slate-400">
            Perbarui informasi dokumen yang sudah ada
        </p>
    </div>

    {#if loadingData}
        <div class="flex justify-center items-center py-20">
            <div
                class="animate-spin rounded-full h-16 w-16 border-4 border-primary border-t-transparent"
            ></div>
        </div>
    {:else}
        <!-- Content -->
        {#if error}
            <div
                class="bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-xl p-4 text-red-700 dark:text-red-400"
            >
                <p class="font-semibold">Error:</p>
                <p>{error}</p>
            </div>
        {/if}

        <div
            class="bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 shadow-sm overflow-hidden"
        >
            <div class="p-6 md:p-8">
                <h3 class="text-lg font-bold mb-6 flex items-center gap-2">
                    <span class="material-symbols-outlined text-primary"
                        >edit_document</span
                    >
                    Detail Informasi Dokumen
                </h3>

                <form class="space-y-6" on:submit|preventDefault={handleSubmit}>
                    <!-- Judul -->
                    <div>
                        <label
                            class="block text-sm font-bold mb-2 text-slate-700 dark:text-slate-300"
                        >
                            Judul Dokumen <span class="text-red-500">*</span>
                        </label>
                        <input
                            bind:value={title}
                            class="w-full px-4 py-2.5 bg-slate-100 dark:bg-slate-800 border-none rounded-lg focus:ring-2 focus:ring-primary/50 text-sm transition-all"
                            placeholder="Contoh: Laporan Keuangan Q3 2023"
                            type="text"
                        />
                    </div>

                    <!-- Penulis + Jenis -->
                    <div class="grid md:grid-cols-2 gap-6">
                        <div>
                            <label
                                class="block text-sm font-bold mb-2 text-slate-700 dark:text-slate-300"
                            >
                                Penulis <span class="text-red-500">*</span>
                            </label>
                            <input
                                bind:value={author}
                                class="w-full px-4 py-2.5 bg-slate-100 dark:bg-slate-800 border-none rounded-lg focus:ring-2 focus:ring-primary/50 text-sm transition-all"
                                placeholder="Masukkan nama penulis"
                                type="text"
                            />
                        </div>

                        <div>
                            <label
                                class="block text-sm font-bold mb-2 text-slate-700 dark:text-slate-300"
                            >
                                Jenis File <span class="text-red-500">*</span>
                            </label>
                            <select
                                bind:value={fileType}
                                class="w-full px-4 py-2.5 bg-slate-100 dark:bg-slate-800 border-none rounded-lg focus:ring-2 focus:ring-primary/50 text-sm transition-all"
                            >
                                <option value="">Pilih jenis file</option>
                                <option value="pdf">PDF</option>
                                <option value="docx">DOCX</option>
                                <option value="xlsx">XLSX</option>
                                <option value="pptx">PPTX</option>
                                <option value="zip">ZIP</option>
                            </select>
                        </div>
                    </div>

                    <!-- Status -->
                    <div>
                        <label
                            class="block text-sm font-bold mb-2 text-slate-700 dark:text-slate-300"
                        >
                            Status <span class="text-red-500">*</span>
                        </label>
                        <select
                            bind:value={status}
                            class="w-full px-4 py-2.5 bg-slate-100 dark:bg-slate-800 border-none rounded-lg focus:ring-2 focus:ring-primary/50 text-sm transition-all"
                        >
                            <option value="draft">Draft</option>
                            <option value="publish">Publish</option>
                        </select>
                    </div>

                    <!-- Upload File (Optional) -->
                    <div>
                        <label
                            class="block text-sm font-bold mb-2 text-slate-700 dark:text-slate-300"
                        >
                            Ganti File (Opsional)
                        </label>
                        <div
                            class="mb-3 p-3 bg-blue-50 dark:bg-blue-900/20 text-blue-700 dark:text-blue-300 rounded-lg text-sm flex items-center gap-2"
                        >
                            <span class="material-symbols-outlined text-base"
                                >info</span
                            >
                            {currentFileName}. Biarkan kosong jika tidak ingin
                            mengganti file.
                        </div>
                        <input
                            type="file"
                            on:change={handleFileChange}
                            class="w-full px-4 py-2.5 bg-slate-100 dark:bg-slate-800 border-none rounded-lg focus:ring-2 focus:ring-primary/50 text-sm transition-all file:mr-4 file:py-2 file:px-4 file:rounded-lg file:border-0 file:bg-primary file:text-white file:font-semibold hover:file:bg-primary/90"
                        />
                        {#if file}
                            <p
                                class="mt-2 text-sm text-green-600 flex items-center gap-1"
                            >
                                <span
                                    class="material-symbols-outlined text-base"
                                    >check_circle</span
                                >
                                File baru dipilih: {file.name}
                            </p>
                        {/if}
                    </div>

                    <!-- Confirm -->
                    <div
                        class="flex gap-3 items-start bg-primary/5 dark:bg-primary/10 p-4 rounded-lg border border-primary/20"
                    >
                        <input
                            type="checkbox"
                            bind:checked={confirm}
                            class="mt-0.5 w-5 h-5 text-primary rounded focus:ring-2 focus:ring-primary/50"
                        />
                        <p class="text-sm text-slate-700 dark:text-slate-300">
                            Saya mengonfirmasi bahwa perubahan data ini sudah
                            benar.
                        </p>
                    </div>

                    <!-- Action -->
                    <div
                        class="flex justify-end gap-3 pt-4 border-t border-slate-200 dark:border-slate-800"
                    >
                        <a
                            href="#/documents"
                            use:link
                            class="px-6 py-2.5 bg-slate-100 dark:bg-slate-800 text-slate-700 dark:text-slate-300 rounded-lg font-bold hover:bg-slate-200 dark:hover:bg-slate-700 transition-all"
                        >
                            Batal
                        </a>

                        <button
                            type="submit"
                            class="flex items-center gap-2 px-6 py-2.5 bg-primary text-white font-bold rounded-lg shadow-lg shadow-primary/25 hover:bg-primary/90 transition-all active:scale-95 disabled:opacity-50 disabled:cursor-not-allowed"
                            disabled={loading}
                        >
                            <span class="material-symbols-outlined text-xl"
                                >save</span
                            >
                            <span
                                >{loading
                                    ? "Menyimpan..."
                                    : "Simpan Perubahan"}</span
                            >
                        </button>
                    </div>
                </form>
            </div>
        </div>
    {/if}
</div>
