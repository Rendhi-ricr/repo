<script>
    import { onMount } from "svelte";
    import { getDocuments } from "../../services/documentService";

    let contributors = [];
    let loading = true;

    onMount(() => {
        loadTopContributors();
    });

    async function loadTopContributors() {
        try {
            loading = true;
            const documents = await getDocuments();

            // Group documents by author
            const authorMap = {};
            documents.forEach((doc) => {
                const author = doc.penulis || "Unknown";
                if (!authorMap[author]) {
                    authorMap[author] = {
                        name: author,
                        papers: 0,
                        types: new Set(),
                    };
                }
                authorMap[author].papers++;
                authorMap[author].types.add(doc.jenis_file);
            });

            // Convert to array, sort by paper count, take top 4
            contributors = Object.values(authorMap)
                .sort((a, b) => b.papers - a.papers)
                .slice(0, 4)
                .map((author, index) => ({
                    ...author,
                    institution:
                        Array.from(author.types).join(", ") || "Dokumen",
                    verified: author.papers >= 3,
                    rank: index + 1,
                }));
        } catch (e) {
            console.error("Failed to load contributors:", e);
        } finally {
            loading = false;
        }
    }

    function getInitials(name) {
        if (!name) return "?";
        return name
            .split(" ")
            .map((n) => n[0])
            .join("")
            .substring(0, 2)
            .toUpperCase();
    }

    function getGradient(index) {
        const gradients = [
            "from-amber-500 to-orange-500",
            "from-slate-400 to-slate-500",
            "from-amber-700 to-amber-800",
            "from-primary to-blue-600",
        ];
        return gradients[index] || gradients[3];
    }
</script>

<section class="py-12 lg:py-16 bg-white dark:bg-[#131d27]">
    <div class="container mx-auto max-w-6xl px-4">
        <div class="text-center mb-10">
            <h2 class="text-3xl font-bold text-slate-900 dark:text-white mb-3">
                Kontributor Teratas
            </h2>
            <p class="text-slate-500 dark:text-slate-400 max-w-2xl mx-auto">
                Penulis dengan kontribusi terbanyak dalam repositori.
            </p>
        </div>

        {#if loading}
            <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
                {#each Array(4) as _}
                    <div
                        class="bg-slate-50 dark:bg-surface-highlight p-6 rounded-xl border border-slate-100 dark:border-slate-700/50 flex flex-col items-center"
                    >
                        <div
                            class="w-20 h-20 rounded-full bg-slate-200 dark:bg-slate-700 animate-pulse mb-4"
                        ></div>
                        <div
                            class="h-5 w-32 bg-slate-200 dark:bg-slate-700 rounded animate-pulse mb-2"
                        ></div>
                        <div
                            class="h-4 w-24 bg-slate-100 dark:bg-slate-800 rounded animate-pulse mb-3"
                        ></div>
                        <div
                            class="h-8 w-28 bg-slate-100 dark:bg-slate-800 rounded-full animate-pulse"
                        ></div>
                    </div>
                {/each}
            </div>
        {:else if contributors.length === 0}
            <div class="text-center py-12">
                <span
                    class="material-symbols-outlined text-5xl text-slate-300 dark:text-slate-700 mb-4"
                    >group_off</span
                >
                <p class="text-slate-500">Belum ada kontributor</p>
            </div>
        {:else}
            <div class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-6">
                {#each contributors as contributor, index}
                    <div
                        class="bg-slate-50 dark:bg-surface-highlight p-6 rounded-xl border border-slate-100 dark:border-slate-700/50 flex flex-col items-center text-center hover:-translate-y-1 transition-transform cursor-pointer"
                    >
                        <div class="relative mb-4">
                            <div
                                class="w-20 h-20 rounded-full bg-gradient-to-br {getGradient(
                                    index,
                                )} flex items-center justify-center text-white text-2xl font-bold border-2 border-primary"
                            >
                                {getInitials(contributor.name)}
                            </div>
                            {#if contributor.verified}
                                <div
                                    class="absolute -bottom-1 -right-1 bg-white dark:bg-surface-highlight rounded-full p-1 shadow-sm"
                                >
                                    <span
                                        class="material-symbols-outlined text-yellow-500 text-lg"
                                        >verified</span
                                    >
                                </div>
                            {/if}
                            {#if index < 3}
                                <div
                                    class="absolute -top-1 -left-1 w-6 h-6 rounded-full bg-gradient-to-br {getGradient(
                                        index,
                                    )} flex items-center justify-center text-white text-xs font-bold shadow-md"
                                >
                                    #{contributor.rank}
                                </div>
                            {/if}
                        </div>

                        <h3
                            class="font-bold text-slate-900 dark:text-white text-lg"
                        >
                            {contributor.name}
                        </h3>
                        <p
                            class="text-primary text-sm font-medium mb-3 line-clamp-1"
                        >
                            {contributor.institution}
                        </p>

                        <div
                            class="flex gap-4 text-xs text-slate-500 dark:text-slate-400 bg-white dark:bg-slate-800 py-2 px-4 rounded-full"
                        >
                            <span
                                ><strong class="text-slate-800 dark:text-white"
                                    >{contributor.papers}</strong
                                > Dokumen</span
                            >
                        </div>
                    </div>
                {/each}
            </div>
        {/if}
    </div>
</section>

<style>
    .bg-surface-highlight {
        background-color: #233648;
    }
    .line-clamp-1 {
        display: -webkit-box;
        -webkit-line-clamp: 1;
        line-clamp: 1;
        -webkit-box-orient: vertical;
        overflow: hidden;
    }
</style>
