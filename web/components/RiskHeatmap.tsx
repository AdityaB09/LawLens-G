"use client";

import {
  ResponsiveContainer,
  BarChart,
  Bar,
  XAxis,
  YAxis,
  Tooltip,
  Legend,
} from "recharts";

type Clause = {
  id: number;
  clauseType: string;
  riskLevel: string;
};

export default function RiskHeatmap({ clauses }: { clauses: Clause[] }) {
  const byType: Record<
    string,
    { type: string; low: number; medium: number; high: number }
  > = {};

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
      <div className="flex items-center justify-between mb-2">
        <h3 className="text-sm font-semibold">Risk Heatmap by Clause Type</h3>
        <span className="text-[10px] text-slate-400">
          Stacked counts of LOW / MEDIUM / HIGH risk clauses
        </span>
      </div>
      <div className="h-64">
        <ResponsiveContainer width="100%" height="100%">
          <BarChart data={data} stackOffset="none">
            <XAxis dataKey="type" tick={{ fontSize: 10, fill: "#cbd5f5" }} />
            <YAxis tick={{ fontSize: 10, fill: "#cbd5f5" }} />
            <Tooltip
              contentStyle={{
                backgroundColor: "#020617",
                border: "1px solid #1e293b",
                fontSize: 11,
              }}
            />
            <Legend
              wrapperStyle={{ fontSize: 11 }}
              formatter={(value) => (
                <span style={{ color: "#e2e8f0" }}>{value.toUpperCase()}</span>
              )}
            />
            {/* soft green for LOW */}
            <Bar dataKey="low" stackId="a" name="Low" fill="#22c55e" />
            {/* amber for MEDIUM */}
            <Bar dataKey="medium" stackId="a" name="Medium" fill="#f59e0b" />
            {/* bright red for HIGH */}
            <Bar dataKey="high" stackId="a" name="High" fill="#ef4444" />
          </BarChart>
        </ResponsiveContainer>
      </div>
    </div>
  );
}
