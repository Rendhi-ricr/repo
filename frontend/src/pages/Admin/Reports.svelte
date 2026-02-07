<script>
    import { onMount } from "svelte";
    import { getDocuments } from "../../services/documentService";

    let loading = true;
    let documents = [];

    // Report data
    let totalDocs = 0;
    let publishedDocs = 0;
    let draftDocs = 0;
    let docsByType = {};
    let docsByMonth = [];

    // Filter
    let selectedPeriod = "all";
    let selectedType = "all";

    onMount(() => {
        loadData();
    });

    async function loadData() {
        try {
            loading = true;
            documents = await getDocuments();
            calculateStats();
        } catch (e) {
            console.error("Failed to load data:", e);
        } finally {
            loading = false;
        }
    }

    function calculateStats() {
        totalDocs = documents.length;
        publishedDocs = documents.filter((d) => d.status === "publish").length;
        draftDocs = documents.filter((d) => d.status === "draft").length;

        // Group by type
        docsByType = documents.reduce((acc, doc) => {
            const type = doc.jenis_file.toUpperCase();
            acc[type] = (acc[type] || 0) + 1;
            return acc;
        }, {});

        // Group by month (last 6 months)
        const months = [];
        for (let i = 5; i >= 0; i--) {
            const date = new Date();
            date.setMonth(date.getMonth() - i);
            const monthName = date.toLocaleDateString("id-ID", {
                month: "short",
            });
            const year = date.getFullYear();
            const monthKey = `${date.getFullYear()}-${String(date.getMonth() + 1).padStart(2, "0")}`;

            const count = documents.filter((doc) => {
                const docDate = new Date(doc.created_at);
                return (
                    docDate.getFullYear() === date.getFullYear() &&
                    docDate.getMonth() === date.getMonth()
                );
            }).length;

            months.push({ name: monthName, year, count, key: monthKey });
        }
        docsByMonth = months;
    }

    function getMaxCount() {
        return Math.max(...docsByMonth.map((m) => m.count), 1);
    }

    function getTypeColor(type) {
        const colors = {
            PDF: {
                bg: "bg-red-500",
                text: "text-red-500",
                light: "bg-red-100 dark:bg-red-900/30",
            },
            DOCX: {
                bg: "bg-blue-500",
                text: "text-blue-500",
                light: "bg-blue-100 dark:bg-blue-900/30",
            },
            XLSX: {
                bg: "bg-green-500",
                text: "text-green-500",
                light: "bg-green-100 dark:bg-green-900/30",
            },
            PPTX: {
                bg: "bg-orange-500",
                text: "text-orange-500",
                light: "bg-orange-100 dark:bg-orange-900/30",
            },
        };
        return (
            colors[type] || {
                bg: "bg-slate-500",
                text: "text-slate-500",
                light: "bg-slate-100 dark:bg-slate-800",
            }
        );
    }

    function exportReport() {
        alert("Fitur export laporan akan segera tersedia!");
    }

    function printReport() {
        window.print();
    }
</script>

<div class="max-w-7xl mx-auto flex flex-col gap-6">
    <!-- Header -->
    <div class="flex flex-wrap items-end justify-between gap-4">
        <div class="flex flex-col gap-1">
            <h2
                class="text-3xl font-black text-slate-900 dark:text-white tracking-tight"
            >
                Laporan & Statistik
            </h2>
            <p class="text-slate-500 dark:text-slate-400">
                Analisis dan ringkasan data dokumen sistem.
            </p>
        </div>
        <div class="flex gap-3">
            <button
                on:click={printReport}
                class="flex items-center gap-2 px-4 py-2.5 bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 text-slate-700 dark:text-slate-300 font-medium rounded-lg hover:bg-slate-50 dark:hover:bg-slate-700 transition-all"
            >
                <span class="material-symbols-outlined text-xl">print</span>
                <span>Cetak</span>
            </button>
            <button
                on:click={exportReport}
                class="flex items-center gap-2 px-5 py-2.5 bg-primary text-white font-bold rounded-lg shadow-lg shadow-primary/25 hover:bg-primary/90 transition-all active:scale-95"
            >
                <span class="material-symbols-outlined text-xl">download</span>
                <span>Export PDF</span>
            </button>
        </div>
    </div>

    <!-- Filter Bar -->
    <div
        class="bg-white dark:bg-slate-900 p-4 rounded-xl border border-slate-200 dark:border-slate-800 shadow-sm flex flex-wrap gap-4"
    >
        <div class="flex items-center gap-2">
            <span class="material-symbols-outlined text-slate-400"
                >filter_alt</span
            >
            <span class="text-sm font-medium text-slate-600 dark:text-slate-400"
                >Filter:</span
            >
        </div>
        <select
            bind:value={selectedPeriod}
            class="px-4 py-2 bg-slate-100 dark:bg-slate-800 border-none rounded-lg text-sm font-medium"
        >
            <option value="all">Semua Waktu</option>
            <option value="today">Hari Ini</option>
            <option value="week">Minggu Ini</option>
            <option value="month">Bulan Ini</option>
            <option value="year">Tahun Ini</option>
        </select>
        <select
            bind:value={selectedType}
            class="px-4 py-2 bg-slate-100 dark:bg-slate-800 border-none rounded-lg text-sm font-medium"
        >
            <option value="all">Semua Tipe</option>
            <option value="pdf">PDF</option>
            <option value="docx">DOCX</option>
            <option value="xlsx">XLSX</option>
            <option value="pptx">PPTX</option>
        </select>
    </div>

    <!-- Summary Cards -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
        <div
            class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 p-6 shadow-sm"
        >
            <div class="flex items-center justify-between mb-4">
                <span class="text-sm font-medium text-slate-500"
                    >Total Dokumen</span
                >
                <div class="p-2 bg-blue-100 dark:bg-blue-900/30 rounded-lg">
                    <span class="material-symbols-outlined text-blue-600"
                        >folder</span
                    >
                </div>
            </div>
            {#if loading}
                <div
                    class="h-10 w-20 bg-slate-200 dark:bg-slate-700 rounded animate-pulse"
                ></div>
            {:else}
                <div class="text-4xl font-black text-slate-900 dark:text-white">
                    {totalDocs}
                </div>
            {/if}
            <p class="text-sm text-slate-500 mt-2">dokumen tersimpan</p>
        </div>

        <div
            class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 p-6 shadow-sm"
        >
            <div class="flex items-center justify-between mb-4">
                <span class="text-sm font-medium text-slate-500"
                    >Dipublikasi</span
                >
                <div
                    class="p-2 bg-emerald-100 dark:bg-emerald-900/30 rounded-lg"
                >
                    <span class="material-symbols-outlined text-emerald-600"
                        >check_circle</span
                    >
                </div>
            </div>
            {#if loading}
                <div
                    class="h-10 w-20 bg-slate-200 dark:bg-slate-700 rounded animate-pulse"
                ></div>
            {:else}
                <div class="text-4xl font-black text-slate-900 dark:text-white">
                    {publishedDocs}
                </div>
            {/if}
            <p class="text-sm text-slate-500 mt-2">
                {totalDocs > 0
                    ? Math.round((publishedDocs / totalDocs) * 100)
                    : 0}% dari total
            </p>
        </div>

        <div
            class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 p-6 shadow-sm"
        >
            <div class="flex items-center justify-between mb-4">
                <span class="text-sm font-medium text-slate-500">Draft</span>
                <div class="p-2 bg-amber-100 dark:bg-amber-900/30 rounded-lg">
                    <span class="material-symbols-outlined text-amber-600"
                        >edit_note</span
                    >
                </div>
            </div>
            {#if loading}
                <div
                    class="h-10 w-20 bg-slate-200 dark:bg-slate-700 rounded animate-pulse"
                ></div>
            {:else}
                <div class="text-4xl font-black text-slate-900 dark:text-white">
                    {draftDocs}
                </div>
            {/if}
            <p class="text-sm text-slate-500 mt-2">menunggu publikasi</p>
        </div>
    </div>

    <!-- Charts Row -->
    <div class="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <!-- Monthly Chart -->
        <div
            class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 p-6 shadow-sm"
        >
            <h3 class="text-lg font-bold text-slate-900 dark:text-white mb-6">
                Tren Upload Dokumen
            </h3>
            <div class="flex items-end gap-4 h-48">
                {#if loading}
                    {#each Array(6) as _}
                        <div class="flex-1 flex flex-col items-center gap-2">
                            <div
                                class="w-full bg-slate-200 dark:bg-slate-700 rounded-t animate-pulse"
                                style="height: 60%"
                            ></div>
                            <div
                                class="h-4 w-8 bg-slate-200 dark:bg-slate-700 rounded animate-pulse"
                            ></div>
                        </div>
                    {/each}
                {:else}
                    {#each docsByMonth as month}
                        <div class="flex-1 flex flex-col items-center gap-2">
                            <div class="w-full flex flex-col justify-end h-40">
                                <div
                                    class="w-full bg-gradient-to-t from-primary to-blue-400 rounded-t transition-all hover:from-primary/90 hover:to-blue-400/90"
                                    style="height: {(month.count /
                                        getMaxCount()) *
                                        100}%; min-height: 4px;"
                                ></div>
                            </div>
                            <span class="text-xs font-medium text-slate-500"
                                >{month.name}</span
                            >
                            <span
                                class="text-sm font-bold text-slate-900 dark:text-white"
                                >{month.count}</span
                            >
                        </div>
                    {/each}
                {/if}
            </div>
        </div>

        <!-- Document Types -->
        <div
            class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 p-6 shadow-sm"
        >
            <h3 class="text-lg font-bold text-slate-900 dark:text-white mb-6">
                Distribusi Tipe File
            </h3>
            {#if loading}
                <div class="space-y-4">
                    {#each Array(4) as _}
                        <div class="flex items-center gap-4">
                            <div
                                class="w-12 h-12 bg-slate-200 dark:bg-slate-700 rounded-lg animate-pulse"
                            ></div>
                            <div
                                class="flex-1 h-4 bg-slate-200 dark:bg-slate-700 rounded animate-pulse"
                            ></div>
                        </div>
                    {/each}
                </div>
            {:else if Object.keys(docsByType).length === 0}
                <div
                    class="flex flex-col items-center justify-center h-48 text-slate-400"
                >
                    <span class="material-symbols-outlined text-4xl mb-2"
                        >pie_chart</span
                    >
                    <p>Belum ada data</p>
                </div>
            {:else}
                <div class="space-y-4">
                    {#each Object.entries(docsByType) as [type, count]}
                        <div class="flex items-center gap-4">
                            <div
                                class="{getTypeColor(type)
                                    .light} p-3 rounded-lg"
                            >
                                <span
                                    class="material-symbols-outlined {getTypeColor(
                                        type,
                                    ).text}">description</span
                                >
                            </div>
                            <div class="flex-1">
                                <div class="flex justify-between mb-1">
                                    <span
                                        class="font-semibold text-slate-900 dark:text-white"
                                        >{type}</span
                                    >
                                    <span class="text-sm text-slate-500"
                                        >{count} file</span
                                    >
                                </div>
                                <div
                                    class="h-2 bg-slate-100 dark:bg-slate-800 rounded-full overflow-hidden"
                                >
                                    <div
                                        class="{getTypeColor(type)
                                            .bg} h-full rounded-full transition-all"
                                        style="width: {(count / totalDocs) *
                                            100}%"
                                    ></div>
                                </div>
                            </div>
                            <span
                                class="text-sm font-bold text-slate-600 dark:text-slate-400"
                            >
                                {Math.round((count / totalDocs) * 100)}%
                            </span>
                        </div>
                    {/each}
                </div>
            {/if}
        </div>
    </div>

    <!-- Activity Log -->
    <div
        class="bg-white dark:bg-slate-900 rounded-2xl border border-slate-200 dark:border-slate-800 shadow-sm overflow-hidden"
    >
        <div class="p-6 border-b border-slate-200 dark:border-slate-800">
            <h3 class="text-lg font-bold text-slate-900 dark:text-white">
                Log Aktivitas Terbaru
            </h3>
            <p class="text-sm text-slate-500">
                Riwayat aktivitas sistem dalam 7 hari terakhir
            </p>
        </div>
        <div class="divide-y divide-slate-100 dark:divide-slate-800">
            {#if loading}
                {#each Array(5) as _}
                    <div class="p-4 flex items-center gap-4">
                        <div
                            class="w-10 h-10 bg-slate-200 dark:bg-slate-700 rounded-full animate-pulse"
                        ></div>
                        <div class="flex-1 space-y-2">
                            <div
                                class="h-4 w-48 bg-slate-200 dark:bg-slate-700 rounded animate-pulse"
                            ></div>
                            <div
                                class="h-3 w-32 bg-slate-100 dark:bg-slate-800 rounded animate-pulse"
                            ></div>
                        </div>
                    </div>
                {/each}
            {:else}
                <div
                    class="p-4 flex items-center gap-4 hover:bg-slate-50 dark:hover:bg-slate-800/50"
                >
                    <div
                        class="w-10 h-10 rounded-full bg-emerald-100 dark:bg-emerald-900/30 flex items-center justify-center"
                    >
                        <span class="material-symbols-outlined text-emerald-600"
                            >upload_file</span
                        >
                    </div>
                    <div class="flex-1">
                        <p class="font-medium text-slate-900 dark:text-white">
                            Dokumen baru diupload
                        </p>
                        <p class="text-sm text-slate-500">
                            Admin mengupload 3 dokumen baru
                        </p>
                    </div>
                    <span class="text-xs text-slate-400">2 jam lalu</span>
                </div>
                <div
                    class="p-4 flex items-center gap-4 hover:bg-slate-50 dark:hover:bg-slate-800/50"
                >
                    <div
                        class="w-10 h-10 rounded-full bg-blue-100 dark:bg-blue-900/30 flex items-center justify-center"
                    >
                        <span class="material-symbols-outlined text-blue-600"
                            >edit</span
                        >
                    </div>
                    <div class="flex-1">
                        <p class="font-medium text-slate-900 dark:text-white">
                            Dokumen diperbarui
                        </p>
                        <p class="text-sm text-slate-500">
                            Metadata dokumen telah diubah
                        </p>
                    </div>
                    <span class="text-xs text-slate-400">5 jam lalu</span>
                </div>
                <div
                    class="p-4 flex items-center gap-4 hover:bg-slate-50 dark:hover:bg-slate-800/50"
                >
                    <div
                        class="w-10 h-10 rounded-full bg-purple-100 dark:bg-purple-900/30 flex items-center justify-center"
                    >
                        <span class="material-symbols-outlined text-purple-600"
                            >person_add</span
                        >
                    </div>
                    <div class="flex-1">
                        <p class="font-medium text-slate-900 dark:text-white">
                            User baru terdaftar
                        </p>
                        <p class="text-sm text-slate-500">
                            2 pengguna baru bergabung
                        </p>
                    </div>
                    <span class="text-xs text-slate-400">Kemarin</span>
                </div>
                <div
                    class="p-4 flex items-center gap-4 hover:bg-slate-50 dark:hover:bg-slate-800/50"
                >
                    <div
                        class="w-10 h-10 rounded-full bg-amber-100 dark:bg-amber-900/30 flex items-center justify-center"
                    >
                        <span class="material-symbols-outlined text-amber-600"
                            >publish</span
                        >
                    </div>
                    <div class="flex-1">
                        <p class="font-medium text-slate-900 dark:text-white">
                            Dokumen dipublikasikan
                        </p>
                        <p class="text-sm text-slate-500">
                            5 dokumen mengubah status ke published
                        </p>
                    </div>
                    <span class="text-xs text-slate-400">2 hari lalu</span>
                </div>
                <div
                    class="p-4 flex items-center gap-4 hover:bg-slate-50 dark:hover:bg-slate-800/50"
                >
                    <div
                        class="w-10 h-10 rounded-full bg-red-100 dark:bg-red-900/30 flex items-center justify-center"
                    >
                        <span class="material-symbols-outlined text-red-600"
                            >delete</span
                        >
                    </div>
                    <div class="flex-1">
                        <p class="font-medium text-slate-900 dark:text-white">
                            Dokumen dihapus
                        </p>
                        <p class="text-sm text-slate-500">
                            1 dokumen dihapus dari sistem
                        </p>
                    </div>
                    <span class="text-xs text-slate-400">3 hari lalu</span>
                </div>
            {/if}
        </div>
    </div>
</div>
