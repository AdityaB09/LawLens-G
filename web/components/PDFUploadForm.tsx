"use client";

import { FormEvent, useState } from "react";
import { useRouter } from "next/navigation";

const API_BASE = process.env.NEXT_PUBLIC_API_BASE || "http://localhost:8080/api";

export default function PDFUploadForm() {
  const router = useRouter();
  const [title, setTitle] = useState("");
  const [partyA, setPartyA] = useState("");
  const [partyB, setPartyB] = useState("");
  const [file, setFile] = useState<File | null>(null);
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  async function onSubmit(e: FormEvent) {
    e.preventDefault();
    setError(null);

    if (!title || !partyA || !partyB || !file) {
      setError("Please fill all fields and choose a PDF file.");
      return;
    }

    const formData = new FormData();
    formData.append("title", title);
    formData.append("partyA", partyA);
    formData.append("partyB", partyB);
    formData.append("file", file);

    setLoading(true);
    try {
      const res = await fetch(`${API_BASE}/contracts/pdf`, {
        method: "POST",
        body: formData
      });

      if (!res.ok) {
        const data = await res.json().catch(() => ({}));
        throw new Error(data.error || "Upload failed");
      }

      const data = await res.json();
      router.push(`/contracts/${data.id}`);
    } catch (err: any) {
      setError(err.message || "Something went wrong");
    } finally {
      setLoading(false);
    }
  }

  return (
    <form
      onSubmit={onSubmit}
      className="space-y-3 rounded-2xl border border-slate-800 bg-slate-900/80 p-4"
    >
      <div className="flex items-center justify-between">
        <h3 className="text-sm font-semibold">Upload Contract (PDF)</h3>
        <span className="text-[10px] text-slate-400">
          Extracts text and analyzes automatically
        </span>
      </div>

      <div className="grid gap-3 md:grid-cols-3 text-xs">
        <div className="flex flex-col gap-1">
          <label className="text-slate-400">Title</label>
          <input
            className="rounded-lg border border-slate-700 bg-slate-950 px-2 py-1"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            placeholder="e.g. Hosting Agreement"
          />
        </div>
        <div className="flex flex-col gap-1">
          <label className="text-slate-400">Party A</label>
          <input
            className="rounded-lg border border-slate-700 bg-slate-950 px-2 py-1"
            value={partyA}
            onChange={(e) => setPartyA(e.target.value)}
            placeholder="Customer Corp"
          />
        </div>
        <div className="flex flex-col gap-1">
          <label className="text-slate-400">Party B</label>
          <input
            className="rounded-lg border border-slate-700 bg-slate-950 px-2 py-1"
            value={partyB}
            onChange={(e) => setPartyB(e.target.value)}
            placeholder="Vendor X"
          />
        </div>
      </div>

      <div className="flex flex-col gap-1 text-xs">
        <label className="text-slate-400">PDF File</label>
        <input
          type="file"
          accept="application/pdf"
          className="text-xs"
          onChange={(e) => setFile(e.target.files?.[0] || null)}
        />
      </div>

      {error && (
        <p className="text-[11px] text-red-400">
          {error}
        </p>
      )}

      <button
        type="submit"
        disabled={loading}
        className="rounded-xl bg-cyan-500 px-4 py-1.5 text-xs font-medium text-slate-950 hover:bg-cyan-400 disabled:opacity-50"
      >
        {loading ? "Uploading..." : "Analyze PDF"}
      </button>
    </form>
  );
}
