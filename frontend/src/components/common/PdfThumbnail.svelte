<script>
    import { onMount } from "svelte";
    import * as pdfjsLib from "pdfjs-dist";

    // Worker lokal
    pdfjsLib.GlobalWorkerOptions.workerSrc = "/pdf.worker.min.mjs";

    export let url = "";
    export let scale = null; // optional

    let canvas;
    let loading = true;
    let error = null;

    // Reactive: re-render jika URL berubah
    $: if (url && canvas) {
        renderPdf();
    }

    async function renderPdf() {
        if (!url || !canvas) return;

        try {
            loading = true;
            error = null;

            // Fetch the PDF as ArrayBuffer first to avoid browser download
            // when the server sets Content-Disposition: attachment.
            const fullUrl = url.startsWith("http")
                ? url
                : `http://localhost:8080${url}`;
            console.log("Fetching PDF from:", fullUrl);

            const resp = await fetch(fullUrl);
            if (!resp.ok) {
                throw new Error(
                    `Network response was not ok: ${resp.status} ${resp.statusText}`,
                );
            }

            const arrayBuffer = await resp.arrayBuffer();
            console.log(
                "PDF ArrayBuffer size:",
                arrayBuffer.byteLength,
                "bytes",
            );

            const loadingTask = pdfjsLib.getDocument({ data: arrayBuffer });
            const pdf = await loadingTask.promise;
            const page = await pdf.getPage(1);

            let finalScale;
            if (scale) {
                finalScale = scale;
            } else {
                // Auto calc for thumbnail (width ~300px)
                const viewportBasic = page.getViewport({ scale: 1.0 });
                // limit max scale agar tidak terlalu pecah utk icon kecil
                finalScale = 300 / viewportBasic.width;
            }

            const viewport = page.getViewport({ scale: finalScale });

            canvas.height = viewport.height;
            canvas.width = viewport.width;

            const renderContext = {
                canvasContext: canvas.getContext("2d"),
                viewport: viewport,
            };

            await page.render(renderContext).promise;
            loading = false;
        } catch (e) {
            console.error("Error rendering PDF:", e);
            loading = false;
            error = "Gagal memuat preview";
        }
    }

    onMount(() => {
        // Initial render handled by reactive statement or manual call if needed
        if (url && canvas) renderPdf();
    });
</script>

<div
    class="w-full h-full flex items-center justify-center bg-white relative overflow-hidden"
>
    {#if loading}
        <div
            class="absolute inset-0 flex items-center justify-center bg-slate-50 z-10"
        >
            <div
                class="animate-spin rounded-full h-6 w-6 border-2 border-primary border-t-transparent"
            ></div>
        </div>
    {/if}

    <canvas
        bind:this={canvas}
        class="max-w-full max-h-full object-contain shadow-sm {loading
            ? 'opacity-0'
            : 'opacity-100'} transition-opacity duration-300"
    ></canvas>
</div>
