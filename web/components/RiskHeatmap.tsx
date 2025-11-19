"use client";

import { ResponsiveContainer, BarChart, Bar, XAxis, YAxis, Tooltip } from "recharts";

type Clause = {
  id: number;
  clauseType: string;
  riskLevel: string;
};

export default function RiskHeatmap({ clauses }: { clauses: Clause[] }) {
  const byType: Record<string, { type: string; low: number; medium: number; high: number }> = {};

  clauses.forEach((cl) => {
    const key = cl.clauseType || "OTHER";
    if (!byType[key]) {
      byType[key] = { type: key, low: 0, medium: 0, high: 0 };
    }
    if (cl.riskLevel === "HIGH") byType[key].high++;
    else if (cl.riskLevel === "MEDIUM") byType[key].medium++;
    else byType[key].low++;
  });

  const data = Object.values(byType);

  return (
    <div className="rounded-2xl border border-slate-800 bg-slate-900/80 p-4">
      <h3 className="text-sm font-semibold mb-3">Risk Heatmap by Clause Type</h3>
      <div className="h-64">
        <ResponsiveContainer width="100%" height="100%">
          <BarChart data={data} stackOffset="none">
            <XAxis dataKey="type" tick={{ fontSize: 10 }} />
            <YAxis tick={{ fontSize: 10 }} />
            <Tooltip />
            <Bar dataKey="low" stackId="a" />
            <Bar dataKey="medium" stackId="a" />
            <Bar dataKey="high" stackId="a" />
          </BarChart>
        </ResponsiveContainer>
      </div>
    </div>
  );
}
