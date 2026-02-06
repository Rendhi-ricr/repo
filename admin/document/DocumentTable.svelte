<script>
  export let documents = [];
  export let onDelete;
</script>

<div class="bg-white dark:bg-slate-900 rounded-xl border border-slate-200 dark:border-slate-800 shadow-sm overflow-hidden">
  <table class="w-full">
    <thead class="bg-slate-50 dark:bg-slate-800/50">
      <tr>
        <th class="px-6 py-4 text-xs font-bold">Judul</th>
        <th class="px-6 py-4 text-xs font-bold">Author</th>
        <th class="px-6 py-4 text-xs font-bold">Kategori</th>
        <th class="px-6 py-4 text-xs font-bold text-right">Aksi</th>
      </tr>
    </thead>

    <tbody>
      {#each documents as doc}
        <tr class="border-t hover:bg-slate-50">
          <td class="px-6 py-4 font-semibold">{doc.title}</td>
          <td class="px-6 py-4">{doc.author}</td>
          <td class="px-6 py-4">{doc.category}</td>
          <td class="px-6 py-4 text-right">
            <button
              class="text-red-500"
              on:click={() => onDelete(doc)}
            >
              Hapus
            </button>
          </td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>
