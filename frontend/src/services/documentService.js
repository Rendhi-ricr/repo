const API = "http://localhost:8080/api/documents";

export const getDocuments = async () => {
  const response = await fetch(API);
  if (!response.ok) throw new Error("Gagal mengambil data dokumen");
  return response.json();
};

export const getDocumentById = async (id) => {
  const response = await fetch(`${API}/${id}`);
  if (!response.ok) throw new Error("Gagal mengambil data dokumen");
  return response.json();
};

export const createDocument = async (formData) => {
  const response = await fetch(API, {
    method: "POST",
    body: formData,
  });
  if (!response.ok) throw new Error("Gagal membuat dokumen");
  return response.json();
};

export const updateDocument = async (id, formData) => {
  const response = await fetch(`${API}/${id}`, {
    method: "PUT",
    body: formData,
  });
  if (!response.ok) throw new Error("Gagal mengupdate dokumen");
  return response.json();
};

export const deleteDocument = async (id) => {
  const response = await fetch(`${API}/${id}`, {
    method: "DELETE",
  });
  if (!response.ok) throw new Error("Gagal menghapus dokumen");
  return response.json();
};

export const downloadDocument = (id) => {
  window.open(`http://localhost:8080/download/${id}`, "_blank");
};

export const getDocumentPages = async (id) => {
  const response = await fetch(`${API}/pages/${id}`);
  if (!response.ok) throw new Error("Gagal mengambil halaman dokumen");
  return response.json();
};
