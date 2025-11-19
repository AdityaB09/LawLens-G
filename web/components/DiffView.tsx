"use client";

import { FormEvent, useState } from "react";
import { compareContracts } from "../lib/api";

type Clause = {
  id: number;
  orderIndex: number;
  heading: string;
  text: string;
  riskLevel: string;
  riskScore: number;
  status?: string;
  riskDelta?: number;
};

type DiffResult = {
  contractAId: number;
  contractBId: number;
  clausesA: Clause[];
  clausesB: Clause[];
};

export default function DiffView() {
  const [aId, setAId] = useState("");
  const [bId, setBId] = useState("");
  const [result, setResult] = useState<DiffResult | null>(null);
  const [loading, setLoading] = useState(false);

  async function onSubmit(e: FormEvent) {
    e.preventDefault();
    if (!aId || !bId) return;
    setLoading(true);
    try {
      const diff = await compareContracts(Number(aId), Number(bId));
      setResult(diff);
    } finally {
      setLoading(false);
    }
  }

  return (
    <div className="space-y-4">
      <form
        onSubmit={onSubmit}
        className="flex flex-wrap gap-3 items-end rounded-2xl border border-slate-800 bg-slate-900/80 p-4"
      >
        <div className="flex flex-col text-xs">
          <label className="mb-1 text-slate-400">Contract A ID</label>
          <input
            className="rounded-lg bg-slate-950 border border-slate-700 px-2 py-1 text-xs"
            value={aId}
            onChange={(e) => setAId(e.target.value)}
            placeholder="e.g. 1"
          />
        </div>
        <div className="flex flex-col text-xs">
          <label className="mb-1 text-slate-400">Contract B ID</label>
          <input
            className="rounded-lg bg-slate-950 border border-slate-700 px-2 py-1 text-xs"
            value={bId}
            onChange={(e) => setBId(e.target.value)}
            placeholder="e.g. 2"
          />
        </div>
        <button
          type="submit"
          className="rounded-xl bg-cyan-500 px-4 py-1.5 text-xs font-medium text-slate-950 hover:bg-cyan-400 disabled:opacity-50"
          disabled={loading}
        >
          {loading ? "Comparing..." : "Compare"}
        </button>
      </form>

      {result && (
        <div className="grid gap-4 md:grid-cols-2">
          <DiffColumn title={`Contract A #${result.contractAId}`} clauses={result.clausesA} />
          <DiffColumn title={`Contract B #${result.contractBId}`} clauses={result.clausesB} />
        </div>
      )}
    </div>
  );
}

function DiffColumn({ title, clauses }: { title: string; clauses: Clause[] }) {
  return (
    <div className="rounded-2xl border border-slate-800 bg-slate-900/80 p-4 space-y-3 max-h-[70vh] overflow-y-auto">
      <h3 className="text-sm font-semibold mb-2">{title}</h3>
      {clauses.map((cl) => {
        let border = "border-slate-700";
        if (cl.status === "ADDED") border = "border-emerald-500/70";
        if (cl.status === "REMOVED") border = "border-red-500/70";
        if (cl.status === "MODIFIED") border = "border-amber-500/70";

        return (
          <div
            key={cl.id}
            className={`rounded-xl border ${border} bg-slate-950/80 p-3 text-xs space-y-1`}
          >
            <div className="flex items-center justify-between">
              <div className="font-semibold">
                #{cl.orderIndex} {cl.heading || ""}
              </div>
              <span className="text-[10px] text-slate-400">{cl.status}</span>
            </div>
            <p className="line-clamp-4 text-slate-100/90">{cl.text}</p>
            {typeof cl.riskDelta === "number" && cl.status === "MODIFIED" && (
              <p className="text-[10px] text-slate-400">
                Risk delta: {(cl.riskDelta * 100).toFixed(0)}%
              </p>
            )}
          </div>
        );
      })}
    </div>
  );
}
