<script>
	let judul = "";
	let penulis = "";
	let jenis_file = "";
	let file = null;
  
  let documents = [];

async function loadDocuments() {
	const res = await fetch("http://localhost:8080/documents");
	documents = await res.json();
}

loadDocuments();


	async function uploadDocument() {
		if (!file) {
			alert("File belum dipilih");
			return;
		}

		const formData = new FormData();
		formData.append("judul", judul);
		formData.append("penulis", penulis);
		formData.append("jenis_file", jenis_file);
		formData.append("file", file);

		const res = await fetch("http://localhost:8080/uploads", {
			method: "POST",
			body: formData
		});

		const data = await res.json();
		console.log(data);
		alert("Upload berhasil");
	}

  async function publishDocument(id) {
	await fetch(`http://localhost:8080/publish/${id}`, {
		method: "POST"
	});

	// refresh list
	loadDocuments();
}

async function deleteDocument(id) {
	if (!confirm("Yakin ingin menghapus dokumen ini?")) return;

	await fetch(`http://localhost:8080/delete/${id}`, {
		method: "DELETE"
	});

	loadDocuments();
}


</script>

<h2>Upload Dokumen</h2>

<input placeholder="Judul" bind:value={judul} /><br /><br />
<input placeholder="Penulis" bind:value={penulis} /><br /><br />
<input placeholder="Jenis File" bind:value={jenis_file} /><br /><br />

<input
	type="file"
	on:change={(e) => {
		const target = e.target;
		if (target && target.files) {
			file = target.files[0];
		}
	}}
/>


<button on:click={uploadDocument}>Upload</button>

<hr />
<h3>Daftar Dokumen</h3>

{#each documents as doc}
	<div style="margin-bottom:10px">
	<b>{doc.judul}</b><br />
	Penulis: {doc.penulis}<br />
	Status: {doc.status}<br />

	{#if doc.status === "draft"}
		<button on:click={() => publishDocument(doc.id)}>
			Publish
		</button>
	{/if}

  <a href={`http://localhost:8080/download/${doc.id}`}>
	<button>Download</button>
</a>

<button on:click={() => deleteDocument(doc.id)}>
	Delete
</button>

</div>

{/each}

