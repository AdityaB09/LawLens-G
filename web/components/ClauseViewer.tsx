"use client";

import { useState } from "react";

type Trigger = {
  triggerPhrase: string;
  triggerType: string;
  severity: string;
};

type Clause = {
  id: number;
  orderIndex: number;
  heading: string;
  text: string;
  clauseType: string;
  riskLevel: string;
  riskScore: number;
  triggers: Trigger[];
};

export default function ClauseViewer({ clauses }: { clauses: Clause[] }) {
  const [selectedIndex, setSelectedIndex] = useState(0);
  const clause = clauses[selectedIndex];

  if (!clause) return <div className="text-sm text-slate-400">No clauses available.</div>;

  return (
    <div className="rounded-2xl border border-slate-800 bg-slate-900/80 p-4 space-y-3">
      <div className="flex items-center justify-between">
        <div>
          <div className="text-xs text-slate-400">
            Clause {clause.orderIndex} / {clauses.length}
          </div>
          <h3 className="text-sm font-semibold">
            {clause.heading || "Clause text"}
          </h3>
        </div>
        <select
          className="rounded-lg bg-slate-900 border border-slate-700 text-xs px-2 py-1"
          value={selectedIndex}
          onChange={(e) => setSelectedIndex(Number(e.target.value))}
        >
          {clauses.map((cl, idx) => (
            <option key={cl.id} value={idx}>
              #{cl.orderIndex} · {cl.clauseType}
            </option>
          ))}
        </select>
      </div>

      <p className="whitespace-pre-wrap text-sm text-slate-100/90">
        {clause.text}
      </p>

      <div className="flex flex-wrap gap-2 text-[11px]">
        <span className="rounded-full border border-slate-600 px-2 py-0.5">
          Type: {clause.clauseType}
        </span>
        <span className="rounded-full border border-slate-600 px-2 py-0.5">
          Risk: {clause.riskLevel} · {(clause.riskScore * 100).toFixed(0)}%
        </span>
        {clause.triggers?.map((t, i) => (
          <span
            key={i}
            className="rounded-full border border-cyan-500/40 bg-cyan-500/10 px-2 py-0.5"
          >
            {t.triggerType} · {t.severity}
          </span>
        ))}
      </div>
    </div>
  );
}
