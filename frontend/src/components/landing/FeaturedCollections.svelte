<script>
    import { onMount } from "svelte";
    import { getDocuments } from "../../services/documentService";

    let loading = true;
    let categories = [];

    onMount(() => {
        loadCategories();
    });

    async function loadCategories() {
        try {
            loading = true;
            const documents = await getDocuments();

            // Group documents by jenis_file (category)
            const categoryMap = {};
            documents.forEach((doc) => {
                const cat = doc.jenis_file || "Lainnya";
                if (!categoryMap[cat]) {
                    categoryMap[cat] = { count: 0, docs: [] };
                }
                categoryMap[cat].count++;
                categoryMap[cat].docs.push(doc);
            });

            // Convert to array and add styling
            const categoryStyles = {
                PDF: {
                    categoryColor: "bg-red-600/90",
                    hoverColor: "group-hover:text-red-400",
                    image: "https://images.unsplash.com/photo-1456513080510-7bf3a84b82f8?w=600&h=400&fit=crop",
                },
                Skripsi: {
                    categoryColor: "bg-primary/90",
                    hoverColor: "group-hover:text-primary",
                    image: "https://images.unsplash.com/photo-1523050854058-8df90110c9f1?w=600&h=400&fit=crop",
                },
                Tesis: {
                    categoryColor: "bg-purple-600/90",
                    hoverColor: "group-hover:text-purple-400",
                    image: "https://images.unsplash.com/photo-1532012197267-da84d127e765?w=600&h=400&fit=crop",
                },
                Jurnal: {
                    categoryColor: "bg-teal-600/90",
                    hoverColor: "group-hover:text-teal-400",
                    image: "https://images.unsplash.com/photo-1481627834876-b7833e8f5570?w=600&h=400&fit=crop",
                },
                Disertasi: {
                    categoryColor: "bg-amber-600/90",
                    hoverColor: "group-hover:text-amber-400",
                    image: "https://images.unsplash.com/photo-1507003211169-0a1dd7228f2d?w=600&h=400&fit=crop",
                },
            };

            const defaultStyle = {
                categoryColor: "bg-slate-600/90",
                hoverColor: "group-hover:text-slate-400",
                image: "https://images.unsplash.com/photo-1497633762265-9d179a990aa6?w=600&h=400&fit=crop",
            };

            categories = Object.entries(categoryMap)
                .map(([name, data]) => ({
                    title: name,
                    category: name,
                    count: `${data.count} Dokumen`,
                    ...(categoryStyles[name] || defaultStyle),
                }))
                .slice(0, 3); // Show max 3 categories
        } catch (e) {
            console.error("Failed to load categories:", e);
        } finally {
            loading = false;
        }
    }
</script>

<section class="py-12 bg-white dark:bg-[#131d27]">
    <div class="container mx-auto max-w-6xl px-4">
        <div class="flex items-center justify-between mb-8">
            <div>
                <h2
                    class="text-2xl md:text-3xl font-bold text-slate-900 dark:text-white mb-2"
                >
                    Koleksi Unggulan
                </h2>
                <p class="text-slate-500 dark:text-slate-400">
                    Kumpulan dokumen berdasarkan kategori.
                </p>
            </div>
            <a
                class="hidden md:flex items-center gap-1 text-primary font-medium hover:underline"
                href="#/documents"
            >
                Lihat Semua <span class="material-symbols-outlined text-sm"
                    >arrow_forward_ios</span
                >
            </a>
        </div>

        {#if loading}
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
                {#each Array(3) as _}
                    <div
                        class="h-64 md:h-80 bg-slate-200 dark:bg-slate-700 rounded-xl animate-pulse"
                    ></div>
                {/each}
            </div>
        {:else if categories.length === 0}
            <div class="text-center py-12">
                <span
                    class="material-symbols-outlined text-5xl text-slate-300 dark:text-slate-700 mb-4"
                    >folder_off</span
                >
                <p class="text-slate-500">Belum ada dokumen tersedia</p>
            </div>
        {:else}
            <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
                {#each categories as collection}
                    <a
                        href="#/documents?category={encodeURIComponent(
                            collection.title,
                        )}"
                        class="group relative overflow-hidden rounded-xl h-64 md:h-80 cursor-pointer block"
                    >
                        <div
                            class="absolute inset-0 bg-cover bg-center transition-transform duration-500 group-hover:scale-110"
                            style="background-image: url('{collection.image}');"
                        ></div>
                        <div
                            class="absolute inset-0 bg-gradient-to-t from-black/90 via-black/40 to-transparent"
                        ></div>
                        <div class="absolute bottom-0 left-0 p-6 w-full">
                            <span
                                class="inline-block px-3 py-1 {collection.categoryColor} text-white text-xs font-bold rounded-full mb-3 backdrop-blur-sm"
                            >
                                {collection.category}
                            </span>
                            <h3
                                class="text-xl font-bold text-white mb-1 {collection.hoverColor} transition-colors"
                            >
                                {collection.title}
                            </h3>
                            <div
                                class="flex items-center justify-between text-slate-300 text-sm"
                            >
                                <span>{collection.count}</span>
                                <span
                                    class="material-symbols-outlined opacity-0 group-hover:opacity-100 transition-opacity -translate-x-2 group-hover:translate-x-0"
                                    >arrow_right_alt</span
                                >
                            </div>
                        </div>
                    </a>
                {/each}
            </div>
        {/if}

        <div class="md:hidden mt-6 text-center">
            <a
                class="inline-flex items-center gap-1 text-primary font-medium hover:underline"
                href="#/documents"
            >
                Lihat Semua <span class="material-symbols-outlined text-sm"
                    >arrow_forward_ios</span
                >
            </a>
        </div>
    </div>
</section>
