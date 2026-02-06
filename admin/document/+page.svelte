<script>
  import { onMount } from "svelte";
  import DocumentTable from "./DocumentTable.svelte";
  import DeleteConfirm from "./DeleteConfirm.svelte";

  let documents = [];
  let loading = true;

  let showDelete = false;
  let selectedDoc = null;

  async function fetchDocuments() {
    loading = true;
    try {
      const res = await fetch("http://localhost:8080/api/documents");
      documents = await res.json();
    } catch (err) {
      console.error(err);
    } finally {
      loading = false;
    }
  }

  function confirmDelete(doc) {
    selectedDoc = doc;
    showDelete = true;
  }

  async function deleteDocument() {
    await fetch(`http://localhost:8080/api/documents/${selectedDoc.id}`, {
      method: "DELETE"
    });

    showDelete = false;
    selectedDoc = null;
    fetchDocuments();
  }

  onMount(fetchDocuments);
</script>

{#if loading}
  <p class="text-slate-500">Loading...</p>
{:else}
  <DocumentTable
    {documents}
    onDelete={confirmDelete}
  />
{/if}

{#if showDelete}
  <DeleteConfirm
    onConfirm={deleteDocument}
    onCancel={() => showDelete = false}
  />
{/if}