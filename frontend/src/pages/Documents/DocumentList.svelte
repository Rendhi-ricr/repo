<script>
    import { onMount } from "svelte";
    import { link } from "svelte-spa-router";
    import {
        getDocuments,
        deleteDocument,
        downloadDocument,
        getDocumentPages,
    } from "../../services/documentService.js";
    // Import komponen thumbnail
    import PdfThumbnail from "../../components/common/PdfThumbnail.svelte";

    let documents = [];
    let loading = true;
    let error = "";
    let searchQuery = "";
    let filterStatus = "all";

    // Modal state
    let showPagesModal = false;
    let currentPages = [];
    let currentDocId = "";
    let currentDocTitle = "";

    // Preview Detail/Lightbox State
    let showDetailModal = false;
    let detailPageUrl = "";
    let detailPageName = "";

    onMount(async () => {
        await loadDocuments();
    });

    async function loadDocuments() {
        try {
            loading = true;
            error = "";
            documents = await getDocuments();
        } catch (e) {
            error = e.message;
        } finally {
            loading = false;
        }
    }

    async function handleDelete(id) {
        if (!confirm("Apakah Anda yakin ingin menghapus dokumen ini?")) return;

        try {
            await deleteDocument(id);
            await loadDocuments();
        } catch (e) {
            alert("Gagal menghapus dokumen: " + e.message);
        }
    }

    function handleDownload(id) {
        downloadDocument(id);
    }

    async function handleShowPages(doc) {
        if (doc.jenis_file.toLowerCase() !== "pdf") {
            alert("Hanya file PDF yang memiliki fitur pemisahan halaman.");
            return;
        }
        try {
            const pages = await getDocumentPages(doc.id);
            if (!pages || pages.length === 0) {
                alert(
                    "File PDF ini belum diproses atau tidak memiliki halaman yang dapat dipisah. Silakan upload ulang untuk memproses.",
                );
                return;
            }
            currentPages = pages;
            currentDocId = doc.id;
            currentDocTitle = doc.judul;
            showPagesModal = true;
        } catch (e) {
            alert("Gagal memuat halaman: " + e.message);
        }
    }

    function closePagesModal() {
        showPagesModal = false;
        currentPages = [];
    }

    function openDetailPreview(docId, page) {
        detailPageUrl = `/preview/split/${docId}/${page}`;
        detailPageName = page;
        showDetailModal = true;
    }

    function closeDetailPreview() {
        showDetailModal = false;
        detailPageUrl = "";
    }

    $: filteredDocuments = documents.filter((doc) => {
        const matchesSearch =
            doc.judul.toLowerCase().includes(searchQuery.toLowerCase()) ||
            doc.penulis.toLowerCase().includes(searchQuery.toLowerCase());
        const matchesStatus =
            filterStatus === "all" || doc.status === filterStatus;
        return matchesSearch && matchesStatus;
    });

    function formatDate(dateString) {
        const date = new Date(dateString);
        return date.toLocaleDateString("id-ID", {
            day: "numeric",
            month: "short",
            year: "numeric",
            hour: "2-digit",
            minute: "2-digit",
        });
    }

    function getFileIcon(jenisFile) {
        const icons = {
            pdf: {
                icon: "picture_as_pdf",
                iconBg: "bg-red-50 dark:bg-red-900/20",
                iconText: "text-red-600",
                badgeBg: "bg-red-100 dark:bg-red-900/30",
                badgeText: "text-red-700 dark:text-red-400",
            },
            docx: {
                icon: "description",
                iconBg: "bg-blue-50 dark:bg-blue-900/20",
                iconText: "text-blue-600",
                badgeBg: "bg-blue-100 dark:bg-blue-900/30",
                badgeText: "text-blue-700 dark:text-blue-400",
            },
            xlsx: {
                icon: "table_chart",
                iconBg: "bg-green-50 dark:bg-green-900/20",
                iconText: "text-green-600",
                badgeBg: "bg-green-100 dark:bg-green-900/30",
                badgeText: "text-green-700 dark:text-green-400",
            },
            pptx: {
                icon: "slideshow",
                iconBg: "bg-purple-50 dark:bg-purple-900/20",
                iconText: "text-purple-600",
                badgeBg: "bg-purple-100 dark:bg-purple-900/30",
                badgeText: "text-purple-700 dark:text-purple-400",
            },
            zip: {
                icon: "folder_zip",
                iconBg: "bg-yellow-50 dark:bg-yellow-900/20",
                iconText: "text-yellow-600",
                badgeBg: "bg-yellow-100 dark:bg-yellow-900/30",
                badgeText: "text-yellow-700 dark:text-yellow-400",
            },
        };
        return (
            icons[jenisFile] || {
                icon: "description",
                iconBg: "bg-slate-50 dark:bg-slate-800/20",
                iconText: "text-slate-600",
                badgeBg: "bg-slate-100 dark:bg-slate-800/30",
                badgeText: "text-slate-700 dark:text-slate-400",
            }
        );
    }
</script>

<div class="max-w-7xl mx-auto flex flex-col gap-6">
    <!-- Header Actions -->
    <div class="flex flex-wrap items-end justify-between gap-4">
        <div class="flex flex-col gap-1">
            <h2
                class="text-3xl font-black text-slate-900 dark:text-white tracking-tight"
            >
                Kelola Dokumen
            </h2>
            <p class="text-slate-500 dark:text-slate-400">
                Pusat kontrol untuk semua dokumen dan arsip digital perusahaan.
            </p>
        </div>
        <a
            href="/documents/add"
            use:link
            class="flex items-center gap-2 px-5 py-2.5 bg-primary text-white font-bold rounded-lg shadow-lg shadow-primary/25 hover:bg-primary/90 transition-all active:scale-95"
        >
            <span class="material-symbols-outlined text-xl">add</span>
            <span>Tambah Dokumen</span>
        </a>
    </div>

    <!-- Filters & Search -->
    <div
        class="bg-white dark:bg-slate-900 p-4 rounded-xl border border-slate-200 dark:border-slate-800 shadow-sm flex flex-col md:flex-row gap-4 items-center"
    >
        <div class="relative w-full md:w-96">
            <span
                class="material-symbols-outlined absolute left-3 top-1/2 -translate-y-1/2 text-slate-400 text-xl"
                >search</span
            >
            <input
                bind:value={searchQuery}
                class="w-full pl-10 pr-4 py-2 bg-slate-100 dark:bg-slate-800 border-none rounded-lg focus:ring-2 focus:ring-primary/50 text-sm transition-all"
                placeholder="Cari judul, penulis, atau ID..."
                type="text"
            />
        </div>
        <div class="flex gap-2 w-full md:w-auto overflow-x-auto pb-1 md:pb-0">
            <select
                bind:value={filterStatus}
                class="flex items-center gap-2 px-4 py-2 bg-slate-100 dark:bg-slate-800 border-none rounded-lg text-sm font-medium"
            >
                <option value="all">Semua Status</option>
                <option value="draft">Draft</option>
                <option value="publish">Published</option>
            </select>
        </div>
    </div>

    {#if error}
        <div
            class="bg-red-50 dark:bg-red-900/20 border border-red-200 dark:border-red-800 rounded-xl p-4 text-red-700 dark:text-red-400"
        >
            <p class="font-semibold">Error:</p>
            <p>{error}</p>
        </div>
    {/if}

    {#if loading}
        <div class="flex justify-center items-center py-20">
            <div
                class="animate-spin rounded-full h-16 w-16 border-4 border-primary border-t-transparent"
            ></div>
        </div>
    {:else if filteredDocuments.length === 0}
        <div
            class="bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 p-12 text-center"
        >
            <span
                class="material-symbols-outlined text-6xl text-slate-300 dark:text-slate-700 mb-4"
                >folder_off</span
            >
            <h3
                class="text-xl font-bold text-slate-700 dark:text-slate-300 mb-2"
            >
                Tidak ada dokumen
            </h3>
            <p class="text-slate-500 dark:text-slate-400 mb-6">
                {searchQuery || filterStatus !== "all"
                    ? "Tidak ada dokumen yang sesuai dengan filter Anda"
                    : "Mulai dengan menambahkan dokumen baru"}
            </p>
            {#if !searchQuery && filterStatus === "all"}
                <a
                    href="/documents/add"
                    use:link
                    class="inline-flex items-center gap-2 px-5 py-2.5 bg-primary text-white font-bold rounded-lg shadow-lg shadow-primary/25 hover:bg-primary/90 transition-all"
                >
                    <span class="material-symbols-outlined">add</span>
                    <span>Tambah Dokumen Pertama</span>
                </a>
            {/if}
        </div>
    {:else}
        <!-- Data Table -->
        <div
            class="bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 shadow-sm overflow-hidden"
        >
            <div class="overflow-x-auto">
                <table class="w-full text-left border-collapse">
                    <thead>
                        <tr
                            class="bg-slate-50 dark:bg-slate-800/50 border-b border-slate-200 dark:border-slate-800"
                        >
                            <th
                                class="px-6 py-4 text-xs font-bold text-slate-500 dark:text-slate-400 uppercase tracking-wider"
                            >
                                Judul Dokumen
                            </th>
                            <th
                                class="px-6 py-4 text-xs font-bold text-slate-500 dark:text-slate-400 uppercase tracking-wider"
                            >
                                Penulis
                            </th>
                            <th
                                class="px-6 py-4 text-xs font-bold text-slate-500 dark:text-slate-400 uppercase tracking-wider"
                            >
                                Tipe File
                            </th>
                            <th
                                class="px-6 py-4 text-xs font-bold text-slate-500 dark:text-slate-400 uppercase tracking-wider"
                            >
                                Tanggal
                            </th>
                            <th
                                class="px-6 py-4 text-xs font-bold text-slate-500 dark:text-slate-400 uppercase tracking-wider text-right"
                            >
                                Aksi
                            </th>
                        </tr>
                    </thead>
                    <tbody
                        class="divide-y divide-slate-100 dark:divide-slate-800"
                    >
                        {#each filteredDocuments as doc (doc.id)}
                            <tr
                                class="hover:bg-slate-50/80 dark:hover:bg-slate-800/30 transition-colors"
                            >
                                <td class="px-6 py-4">
                                    <div class="flex items-center gap-3">
                                        <div
                                            class="{getFileIcon(doc.jenis_file)
                                                .iconBg} {getFileIcon(
                                                doc.jenis_file,
                                            ).iconText} p-2 rounded-lg"
                                        >
                                            <span
                                                class="material-symbols-outlined"
                                                >{getFileIcon(doc.jenis_file)
                                                    .icon}</span
                                            >
                                        </div>
                                        <div class="flex flex-col">
                                            <span
                                                class="text-sm font-bold text-slate-900 dark:text-white"
                                                >{doc.judul}</span
                                            >
                                            <span class="text-xs text-slate-500"
                                                >ID: {doc.id.substring(
                                                    0,
                                                    8,
                                                )}</span
                                            >
                                        </div>
                                    </div>
                                </td>
                                <td
                                    class="px-6 py-4 text-sm text-slate-600 dark:text-slate-400"
                                >
                                    {doc.penulis}
                                </td>
                                <td class="px-6 py-4">
                                    <span
                                        class="px-2.5 py-1 {getFileIcon(
                                            doc.jenis_file,
                                        ).badgeBg} {getFileIcon(doc.jenis_file)
                                            .badgeText} text-[10px] font-bold uppercase rounded leading-none"
                                    >
                                        {doc.jenis_file}
                                    </span>
                                </td>
                                <td
                                    class="px-6 py-4 text-sm text-slate-600 dark:text-slate-400"
                                >
                                    {formatDate(doc.created_at)}
                                </td>
                                <td class="px-6 py-4 text-right">
                                    <div class="flex justify-end gap-2">
                                        {#if doc.jenis_file.toLowerCase() === "pdf"}
                                            <button
                                                on:click={() =>
                                                    handleShowPages(doc)}
                                                class="p-1.5 text-purple-600 hover:bg-purple-50 dark:hover:bg-purple-900/20 rounded transition-colors"
                                                title="Lihat Halaman"
                                            >
                                                <span
                                                    class="material-symbols-outlined text-xl"
                                                    >menu_book</span
                                                >
                                            </button>
                                        {/if}
                                        <button
                                            on:click={() =>
                                                handleDownload(doc.id)}
                                            class="p-1.5 text-green-600 hover:bg-green-50 dark:hover:bg-green-900/20 rounded transition-colors"
                                            title="Download"
                                        >
                                            <span
                                                class="material-symbols-outlined text-xl"
                                                >download</span
                                            >
                                        </button>
                                        <a
                                            href="/documents/edit/{doc.id}"
                                            use:link
                                            class="p-1.5 text-primary hover:bg-primary/10 rounded transition-colors"
                                            title="Edit"
                                        >
                                            <span
                                                class="material-symbols-outlined text-xl"
                                                >edit_square</span
                                            >
                                        </a>
                                        <button
                                            on:click={() =>
                                                handleDelete(doc.id)}
                                            class="p-1.5 text-slate-400 hover:text-red-500 hover:bg-red-50 dark:hover:bg-red-900/20 rounded transition-colors"
                                            title="Hapus"
                                        >
                                            <span
                                                class="material-symbols-outlined text-xl"
                                                >delete</span
                                            >
                                        </button>
                                    </div>
                                </td>
                            </tr>
                        {/each}
                    </tbody>
                </table>
            </div>

            <!-- Pagination -->
            <div
                class="px-6 py-4 bg-slate-50 dark:bg-slate-800/50 flex items-center justify-between"
            >
                <span class="text-sm text-slate-500">
                    Menampilkan {filteredDocuments.length} dari {documents.length}
                    dokumen
                </span>
                <div class="flex gap-2">
                    <button
                        class="size-8 flex items-center justify-center rounded border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-600 dark:text-slate-400 disabled:opacity-50"
                    >
                        <span class="material-symbols-outlined text-lg"
                            >chevron_left</span
                        >
                    </button>
                    <button
                        class="size-8 flex items-center justify-center rounded border border-primary bg-primary text-white font-bold text-sm"
                    >
                        1
                    </button>
                    <button
                        class="size-8 flex items-center justify-center rounded border border-slate-300 dark:border-slate-700 bg-white dark:bg-slate-900 text-slate-600 dark:text-slate-400"
                    >
                        <span class="material-symbols-outlined text-lg"
                            >chevron_right</span
                        >
                    </button>
                </div>
            </div>
        </div>
    {/if}

    {#if showPagesModal}
        <div
            class="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/50 backdrop-blur-sm"
            on:click={closePagesModal}
        >
            <div
                class="bg-white dark:bg-slate-900 rounded-xl shadow-2xl w-full max-w-4xl max-h-[90vh] flex flex-col overflow-hidden"
                on:click|stopPropagation
            >
                <div
                    class="p-6 border-b border-slate-200 dark:border-slate-800 flex justify-between items-center"
                >
                    <div>
                        <h3
                            class="text-xl font-bold text-slate-900 dark:text-white"
                        >
                            Daftar Halaman
                        </h3>
                        <p class="text-sm text-slate-500">{currentDocTitle}</p>
                    </div>
                    <button
                        on:click={closePagesModal}
                        class="text-slate-400 hover:text-slate-600 dark:hover:text-slate-200"
                    >
                        <span class="material-symbols-outlined">close</span>
                    </button>
                </div>

                <div
                    class="p-6 overflow-y-auto bg-slate-50 dark:bg-slate-900/50"
                >
                    <div
                        class="grid grid-cols-2 md:grid-cols-4 lg:grid-cols-6 gap-4"
                    >
                        {#each currentPages as page}
                            <button
                                on:click={() =>
                                    openDetailPreview(currentDocId, page)}
                                class="flex flex-col gap-2 p-2 bg-white dark:bg-slate-800 rounded-lg border border-slate-200 dark:border-slate-700 hover:border-primary hover:shadow-md transition-all group text-center cursor-pointer"
                            >
                                <div
                                    class="aspect-[3/4] w-full bg-slate-100 dark:bg-slate-700 rounded flex items-center justify-center text-slate-400 group-hover:text-primary relative overflow-hidden shadow-sm"
                                >
                                    <!-- Debug: Matikan sementara PdfThumbnail untuk cek auto-download -->
                                    <PdfThumbnail
                                        url={`/preview/split/${currentDocId}/${page}`}
                                    />

                                    <!-- Transparent layer for clicking -->
                                    <div
                                        class="absolute inset-0 z-10 bg-transparent group-hover:bg-primary/5 transition-colors"
                                    ></div>
                                </div>
                                <span
                                    class="text-xs font-medium truncate px-1 text-slate-600 dark:text-slate-300 group-hover:text-primary w-full"
                                >
                                    {page}
                                </span>
                            </button>
                        {/each}
                    </div>
                </div>
            </div>
        </div>
    {/if}

    <!-- Detail Preview Modal (Lightbox) -->
    {#if showDetailModal}
        <div
            class="fixed inset-0 z-[60] flex items-center justify-center p-4 bg-black/90 backdrop-blur-md"
            on:click={closeDetailPreview}
        >
            <button
                on:click={closeDetailPreview}
                class="absolute top-4 right-4 text-white/50 hover:text-white transition-colors z-[70]"
            >
                <span class="material-symbols-outlined text-4xl">close</span>
            </button>

            <div
                class="w-full h-full flex flex-col items-center justify-center pointer-events-none"
            >
                <div
                    class="h-[85vh] max-w-full bg-white shadow-2xl overflow-auto rounded-lg pointer-events-auto flex items-center justify-center p-4"
                    on:click|stopPropagation
                >
                    <!-- High Res Preview (Scale 2.0) -->
                    <PdfThumbnail url={detailPageUrl} scale={2.0} />
                </div>
                <p class="text-white mt-4 font-medium text-lg drop-shadow-md">
                    {detailPageName}
                </p>
            </div>
        </div>
    {/if}
</div>
