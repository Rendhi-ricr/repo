<script>
  export let onConfirm;
  export let onCancel;
</script>

<div class="fixed inset-0 bg-black/40 flex items-center justify-center z-50">
  <div class="bg-white rounded-lg p-6 w-80">
    <h3 class="font-bold mb-4">Hapus Dokumen?</h3>
    <p class="text-sm text-slate-600 mb-6">
      Data yang dihapus tidak bisa dikembalikan.
    </p>

    <div class="flex justify-end gap-2">
      <button
        class="px-4 py-2 border rounded"
        on:click={onCancel}
      >
        Batal
      </button>
      <button
        class="px-4 py-2 bg-red-600 text-white rounded"
        on:click={onConfirm}
      >
        Hapus
      </button>
    </div>
  </div>
</div>
