"use client";

type Clause = {
  id: number;
  orderIndex: number;
  heading: string;
  clauseType: string;
  riskLevel: string;
};

export default function ClauseList({ clauses }: { clauses: Clause[] }) {
  return (
    <div className="space-y-2 max-h-[70vh] overflow-y-auto pr-1">
      {clauses.map((cl) => {
        const color =
          cl.riskLevel === "HIGH"
            ? "bg-red-500/20 text-red-200 border-red-500/40"
            : cl.riskLevel === "MEDIUM"
            ? "bg-amber-500/20 text-amber-100 border-amber-500/40"
            : "bg-emerald-500/20 text-emerald-100 border-emerald-500/40";

        return (
          <div
            key={cl.id}
            className={`flex items-center justify-between rounded-xl border px-3 py-2 text-xs ${color}`}
          >
            <div>
              <div className="font-medium">
                {cl.orderIndex}. {cl.heading || "Untitled clause"}
              </div>
              <div className="text-[10px] uppercase tracking-wide text-slate-200">
                {cl.clauseType}
              </div>
            </div>
            <span className="text-[10px] opacity-80">{cl.riskLevel}</span>
          </div>
        );
      })}
    </div>
  );
}
