<script>
  let form = {
    title: "",
    author: "",
    file_type: "",
    status: "draft"
  };

  let loading = false;
  let error = "";

  async function submitForm() {
    error = "";
    loading = true;

    try {
      const res = await fetch("http://localhost:8080/api/documents", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify(form)
      });

      if (!res.ok) {
        throw new Error("Gagal menyimpan dokumen");
      }

      alert("Dokumen berhasil ditambahkan");
      
      // reset form
      form = {
        title: "",
        author: "",
        file_type: "",
        status: "draft"
      };
    } catch (err) {
      error = err.message;
    } finally {
      loading = false;
    }
  }
</script>


<form class="space-y-6" on:submit|preventDefault={submitForm}>

  <input
  bind:value={form.title}
  id="judul"
  type="text"
  required
/>
<input
  bind:value={form.author}
  id="penulis"
  type="text"
  required
/>
<select
  bind:value={form.file_type}
  id="jenis_file"
  required
>

  <input bind:value={document.category} placeholder="Kategori" />
  <select bind:value={document.status}>
    <option value="draft">Draft</option>
    <option value="published">Published</option>
  </select>

  <button
  type="submit"
  disabled={loading}
  class="px-8 py-2.5 bg-primary text-white rounded-lg text-sm font-bold"
>
  {#if loading}
    Menyimpan...
  {:else}
    Simpan Dokumen
  {/if}
</button>

</form>
