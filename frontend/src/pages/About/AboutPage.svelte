<script>
    const features = [
        {
            icon: "library_books",
            title: "Koleksi Lengkap",
            description:
                "Akses ribuan karya ilmiah dari berbagai fakultas dan program studi.",
        },
        {
            icon: "search",
            title: "Pencarian Cerdas",
            description:
                "Temukan dokumen dengan mudah menggunakan fitur pencarian yang powerful.",
        },
        {
            icon: "download",
            title: "Unduh Gratis",
            description:
                "Semua dokumen dapat diunduh secara gratis untuk keperluan akademik.",
        },
        {
            icon: "security",
            title: "Aman & Terpercaya",
            description:
                "Dokumen terverifikasi dan tersimpan dengan sistem keamanan tinggi.",
        },
        {
            icon: "devices",
            title: "Akses Multi-Platform",
            description:
                "Buka repositori kapan saja, di mana saja, dari perangkat apa saja.",
        },
        {
            icon: "update",
            title: "Update Berkala",
            description:
                "Koleksi terus bertambah dengan karya ilmiah terbaru setiap harinya.",
        },
    ];

    const team = [
        {
            name: "Administrator",
            role: "Pengelola Sistem",
            icon: "admin_panel_settings",
        },
        {
            name: "Pustakawan",
            role: "Kurator Dokumen",
            icon: "person",
        },
        {
            name: "Tim IT",
            role: "Pengembang Sistem",
            icon: "code",
        },
    ];

    const stats = [
        { value: "2024", label: "Tahun Berdiri" },
        { value: "24/7", label: "Akses Online" },
        { value: "100%", label: "Gratis" },
        { value: "∞", label: "Tanpa Batas" },
    ];
</script>

<div class="min-h-screen bg-slate-50 dark:bg-background-dark">
    <!-- Hero Section -->
    <section
        class="relative overflow-hidden bg-gradient-to-br from-primary via-blue-600 to-indigo-700 text-white py-20 px-4"
    >
        <!-- Decorative Elements -->
        <div
            class="absolute top-0 left-0 w-96 h-96 bg-white/5 rounded-full blur-3xl -translate-x-1/2 -translate-y-1/2"
        ></div>
        <div
            class="absolute bottom-0 right-0 w-96 h-96 bg-blue-400/10 rounded-full blur-3xl translate-x-1/2 translate-y-1/2"
        ></div>

        <div class="container mx-auto max-w-6xl relative z-10">
            <div class="text-center max-w-3xl mx-auto">
                <div
                    class="inline-flex items-center gap-2 px-4 py-2 rounded-full bg-white/10 backdrop-blur-sm border border-white/20 text-sm font-medium mb-6"
                >
                    <span class="material-symbols-outlined text-sm">info</span>
                    Tentang Kami
                </div>

                <h1
                    class="text-4xl md:text-5xl lg:text-6xl font-black mb-6 leading-tight"
                >
                    Repositori <span class="text-blue-200"
                        >Universitas Negeri</span
                    >
                </h1>

                <p class="text-xl text-blue-100 leading-relaxed">
                    Platform digital untuk menyimpan, mengelola, dan
                    menyebarluaskan karya ilmiah civitas akademika Universitas
                    Negeri.
                </p>
            </div>
        </div>
    </section>

    <!-- Stats Section -->
    <section
        class="py-12 bg-white dark:bg-slate-900 border-b border-slate-200 dark:border-slate-800"
    >
        <div class="container mx-auto max-w-6xl px-4">
            <div class="grid grid-cols-2 md:grid-cols-4 gap-6">
                {#each stats as stat}
                    <div class="text-center">
                        <p
                            class="text-3xl md:text-4xl font-black text-primary mb-2"
                        >
                            {stat.value}
                        </p>
                        <p
                            class="text-slate-500 dark:text-slate-400 font-medium"
                        >
                            {stat.label}
                        </p>
                    </div>
                {/each}
            </div>
        </div>
    </section>

    <!-- About Content -->
    <section class="py-16 px-4">
        <div class="container mx-auto max-w-6xl">
            <div class="grid md:grid-cols-2 gap-12 items-center">
                <div>
                    <h2
                        class="text-3xl font-bold text-slate-900 dark:text-white mb-6"
                    >
                        Visi & Misi Kami
                    </h2>

                    <div class="space-y-6">
                        <div
                            class="p-6 rounded-xl bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 shadow-sm"
                        >
                            <div class="flex items-center gap-3 mb-3">
                                <div class="p-2 rounded-lg bg-primary/10">
                                    <span
                                        class="material-symbols-outlined text-primary"
                                        >visibility</span
                                    >
                                </div>
                                <h3
                                    class="text-lg font-bold text-slate-900 dark:text-white"
                                >
                                    Visi
                                </h3>
                            </div>
                            <p class="text-slate-600 dark:text-slate-400">
                                Menjadi pusat informasi dan dokumentasi karya
                                ilmiah terdepan yang mendukung pengembangan ilmu
                                pengetahuan dan teknologi di Indonesia.
                            </p>
                        </div>

                        <div
                            class="p-6 rounded-xl bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700 shadow-sm"
                        >
                            <div class="flex items-center gap-3 mb-3">
                                <div class="p-2 rounded-lg bg-primary/10">
                                    <span
                                        class="material-symbols-outlined text-primary"
                                        >flag</span
                                    >
                                </div>
                                <h3
                                    class="text-lg font-bold text-slate-900 dark:text-white"
                                >
                                    Misi
                                </h3>
                            </div>
                            <ul
                                class="text-slate-600 dark:text-slate-400 space-y-2"
                            >
                                <li class="flex items-start gap-2">
                                    <span
                                        class="material-symbols-outlined text-primary text-sm mt-1"
                                        >check_circle</span
                                    >
                                    Menyediakan akses mudah terhadap karya ilmiah
                                </li>
                                <li class="flex items-start gap-2">
                                    <span
                                        class="material-symbols-outlined text-primary text-sm mt-1"
                                        >check_circle</span
                                    >
                                    Mendokumentasikan hasil penelitian civitas akademika
                                </li>
                                <li class="flex items-start gap-2">
                                    <span
                                        class="material-symbols-outlined text-primary text-sm mt-1"
                                        >check_circle</span
                                    >
                                    Mendukung pengembangan riset dan inovasi
                                </li>
                            </ul>
                        </div>
                    </div>
                </div>

                <div class="relative">
                    <div
                        class="aspect-square rounded-2xl bg-gradient-to-br from-primary/20 to-blue-600/20 p-8 flex items-center justify-center"
                    >
                        <div class="grid grid-cols-2 gap-4 w-full max-w-xs">
                            <div
                                class="aspect-square rounded-xl bg-white dark:bg-slate-800 shadow-lg flex items-center justify-center"
                            >
                                <span
                                    class="material-symbols-outlined text-5xl text-primary"
                                    >school</span
                                >
                            </div>
                            <div
                                class="aspect-square rounded-xl bg-white dark:bg-slate-800 shadow-lg flex items-center justify-center"
                            >
                                <span
                                    class="material-symbols-outlined text-5xl text-blue-500"
                                    >science</span
                                >
                            </div>
                            <div
                                class="aspect-square rounded-xl bg-white dark:bg-slate-800 shadow-lg flex items-center justify-center"
                            >
                                <span
                                    class="material-symbols-outlined text-5xl text-purple-500"
                                    >auto_stories</span
                                >
                            </div>
                            <div
                                class="aspect-square rounded-xl bg-white dark:bg-slate-800 shadow-lg flex items-center justify-center"
                            >
                                <span
                                    class="material-symbols-outlined text-5xl text-teal-500"
                                    >biotech</span
                                >
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </section>

    <!-- Features Section -->
    <section class="py-16 px-4 bg-white dark:bg-slate-900">
        <div class="container mx-auto max-w-6xl">
            <div class="text-center mb-12">
                <h2
                    class="text-3xl font-bold text-slate-900 dark:text-white mb-4"
                >
                    Fitur Unggulan
                </h2>
                <p class="text-slate-500 dark:text-slate-400 max-w-2xl mx-auto">
                    Berbagai kemudahan untuk mengakses dan mengelola karya
                    ilmiah
                </p>
            </div>

            <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
                {#each features as feature}
                    <div
                        class="p-6 rounded-xl bg-slate-50 dark:bg-slate-800 border border-slate-200 dark:border-slate-700 hover:border-primary hover:shadow-lg transition-all group"
                    >
                        <div
                            class="w-14 h-14 rounded-xl bg-primary/10 flex items-center justify-center mb-4 group-hover:bg-primary group-hover:scale-110 transition-all"
                        >
                            <span
                                class="material-symbols-outlined text-2xl text-primary group-hover:text-white"
                                >{feature.icon}</span
                            >
                        </div>
                        <h3
                            class="text-lg font-bold text-slate-900 dark:text-white mb-2"
                        >
                            {feature.title}
                        </h3>
                        <p class="text-slate-500 dark:text-slate-400">
                            {feature.description}
                        </p>
                    </div>
                {/each}
            </div>
        </div>
    </section>

    <!-- Team Section -->
    <section class="py-16 px-4">
        <div class="container mx-auto max-w-6xl">
            <div class="text-center mb-12">
                <h2
                    class="text-3xl font-bold text-slate-900 dark:text-white mb-4"
                >
                    Tim Pengelola
                </h2>
                <p class="text-slate-500 dark:text-slate-400 max-w-2xl mx-auto">
                    Didukung oleh tim profesional yang berdedikasi
                </p>
            </div>

            <div
                class="grid grid-cols-1 md:grid-cols-3 gap-6 max-w-3xl mx-auto"
            >
                {#each team as member}
                    <div
                        class="text-center p-6 rounded-xl bg-white dark:bg-slate-800 border border-slate-200 dark:border-slate-700"
                    >
                        <div
                            class="w-20 h-20 rounded-full bg-gradient-to-br from-primary to-blue-600 flex items-center justify-center mx-auto mb-4"
                        >
                            <span
                                class="material-symbols-outlined text-3xl text-white"
                                >{member.icon}</span
                            >
                        </div>
                        <h3
                            class="text-lg font-bold text-slate-900 dark:text-white"
                        >
                            {member.name}
                        </h3>
                        <p class="text-primary font-medium text-sm">
                            {member.role}
                        </p>
                    </div>
                {/each}
            </div>
        </div>
    </section>

    <!-- Contact Section -->
    <section
        class="py-16 px-4 bg-gradient-to-br from-primary via-blue-600 to-indigo-700 text-white"
    >
        <div class="container mx-auto max-w-6xl">
            <div class="text-center max-w-2xl mx-auto">
                <h2 class="text-3xl font-bold mb-4">Hubungi Kami</h2>
                <p class="text-blue-100 mb-8">
                    Punya pertanyaan atau saran? Jangan ragu untuk menghubungi
                    kami.
                </p>

                <div class="grid grid-cols-1 md:grid-cols-3 gap-6">
                    <div
                        class="p-6 rounded-xl bg-white/10 backdrop-blur-sm border border-white/20"
                    >
                        <span class="material-symbols-outlined text-3xl mb-3"
                            >location_on</span
                        >
                        <h3 class="font-bold mb-1">Alamat</h3>
                        <p class="text-blue-100 text-sm">
                            Jl. Universitas No. 1, Indonesia
                        </p>
                    </div>
                    <div
                        class="p-6 rounded-xl bg-white/10 backdrop-blur-sm border border-white/20"
                    >
                        <span class="material-symbols-outlined text-3xl mb-3"
                            >email</span
                        >
                        <h3 class="font-bold mb-1">Email</h3>
                        <p class="text-blue-100 text-sm">
                            repository@univ.ac.id
                        </p>
                    </div>
                    <div
                        class="p-6 rounded-xl bg-white/10 backdrop-blur-sm border border-white/20"
                    >
                        <span class="material-symbols-outlined text-3xl mb-3"
                            >call</span
                        >
                        <h3 class="font-bold mb-1">Telepon</h3>
                        <p class="text-blue-100 text-sm">(021) 1234567</p>
                    </div>
                </div>
            </div>
        </div>
    </section>

    <!-- Footer -->
    <footer class="py-8 px-4 bg-slate-900 text-slate-400">
        <div class="container mx-auto max-w-6xl text-center">
            <p>© 2024 Repositori Universitas Negeri. All rights reserved.</p>
        </div>
    </footer>
</div>
