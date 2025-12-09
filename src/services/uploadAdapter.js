/*
  uploadAdapter.js
  Exposes uploadFile(file) -> Promise<{ url }>
  Behavior:
  - If REACT_APP_UPLOAD_URL is set, perform a POST to that URL with the file in form-data under 'file'.
    If REACT_APP_UPLOAD_TOKEN exists, add Authorization: Bearer <token> header.
  - Otherwise, fall back to a simulated upload that resolves with a data: URL so you can test locally.

  Note: this keeps real credentials out of source. Use env vars in your .env file:
    REACT_APP_UPLOAD_URL=https://api.example.com/uploads
    REACT_APP_UPLOAD_TOKEN=xxxxx
*/

export async function uploadFile(file) {
  const uploadUrl = process.env.REACT_APP_UPLOAD_URL;
  const token = process.env.REACT_APP_UPLOAD_TOKEN;

  if (uploadUrl) {
    const form = new FormData();
    form.append('file', file);

    const headers = {};
    if (token) headers['Authorization'] = `Bearer ${token}`;

    const resp = await fetch(uploadUrl, {
      method: 'POST',
      body: form,
      headers,
      credentials: 'include',
    });

    if (!resp.ok) {
      const text = await resp.text();
      throw new Error(`Upload failed: ${resp.status} ${text}`);
    }

    // Expecting JSON with { url: '...' }
    const json = await resp.json();
    if (json && (json.url || json.fileUrl || json.file)) {
      return { url: json.url || json.fileUrl || json.file };
    }

    // fallback: return entire body as string under url
    return { url: JSON.stringify(json) };
  }

  // Simulated upload: convert file to data URL and return it
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onload = () => resolve({ url: reader.result });
    reader.onerror = (e) => reject(e);
    reader.readAsDataURL(file);
  });
}
