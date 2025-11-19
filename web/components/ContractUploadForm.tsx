"use client";

import { FormEvent, useState } from "react";
import { useRouter } from "next/navigation";

const API_BASE = process.env.NEXT_PUBLIC_API_BASE || "http://localhost:8080/api";

export default function ContractUploadForm() {
  const router = useRouter();
  const [title, setTitle] = useState("");
  const [partyA, setPartyA] = useState("");
  const [partyB, setPartyB] = useState("");
  const [text, setText] = useState("");
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState<string | null>(null);

  async function onSubmit(e: FormEvent) {
    e.preventDefault();
    setError(null);

    if (!title || !partyA || !partyB || !text) {
      setError("Please fill all fields and paste contract text.");
      return;
    }

    setLoading(true);
    try {
      const res = await fetch(`${API_BASE}/contracts`, {
        method: "POST",
        headers: { "content-type": "application/json" },
        body: JSON.stringify({ title, partyA, partyB, text }),
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
        <h3 className="text-sm font-semibold">Upload Contract (Paste Text)</h3>
        <span className="text-[10px] text-slate-400">
          Creates a new analyzed contract
        </span>
      </div>

      <div className="grid gap-3 md:grid-cols-3 text-xs">
        <div className="flex flex-col gap-1">
          <label className="text-slate-400">Title</label>
          <input
            className="rounded-lg border border-slate-700 bg-slate-950 px-2 py-1"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            placeholder="Master Services Agreement"
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
        <label className="text-slate-400">Contract Text</label>
        <textarea
          className="h-40 rounded-lg border border-slate-700 bg-slate-950 px-2 py-1 text-xs leading-snug"
          value={text}
          onChange={(e) => setText(e.target.value)}
          placeholder="Paste full contract text here..."
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
        {loading ? "Analyzing..." : "Analyze Contract"}
      </button>
    </form>
  );
}
