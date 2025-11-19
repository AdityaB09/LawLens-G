"use client";

type ContractSummary = {
  id: number;
  title: string;
  partyA: string;
  partyB: string;
  uploadedAt: string;
  overallRiskLevel: string;
  overallRiskScore: number;
  clauseCount: number;
  obligationCount: number;
};

export default function ContractCard({ contract }: { contract: ContractSummary }) {
  const riskColor =
    contract.overallRiskLevel === "HIGH"
      ? "bg-red-500/20 text-red-300 border-red-500/50"
      : contract.overallRiskLevel === "MEDIUM"
      ? "bg-amber-500/20 text-amber-300 border-amber-500/50"
      : "bg-emerald-500/20 text-emerald-300 border-emerald-500/50";

  return (
    <a
      href={`/contracts/${contract.id}`}
      className="group rounded-2xl border border-slate-800 bg-slate-900/70 p-4 shadow-lg shadow-slate-950/50 hover:border-cyan-500/60 hover:bg-slate-900"
    >
      <div className="flex items-center justify-between mb-2">
        <h3 className="font-semibold text-sm group-hover:text-cyan-300">
          {contract.title}
        </h3>
        <span className={`rounded-full border px-2 py-0.5 text-[10px] ${riskColor}`}>
          {contract.overallRiskLevel} · {(contract.overallRiskScore * 100).toFixed(0)}%
        </span>
      </div>
      <p className="text-xs text-slate-400 mb-3">
        {contract.partyA} ↔ {contract.partyB}
      </p>
      <div className="flex items-center justify-between text-[11px] text-slate-400">
        <span>{contract.clauseCount} clauses</span>
        <span>{contract.obligationCount} obligations</span>
      </div>
    </a>
  );
}
