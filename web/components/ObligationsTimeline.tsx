"use client";

type Obligation = {
  id: number;
  description: string;
  dueDate?: string | null;
};

export default function ObligationsTimeline({ obligations }: { obligations: Obligation[] }) {
  if (!obligations.length) {
    return (
      <div className="rounded-2xl border border-slate-800 bg-slate-900/80 p-4 text-xs text-slate-400">
        No explicit obligations detected yet.
      </div>
    );
  }

  return (
    <div className="rounded-2xl border border-slate-800 bg-slate-900/80 p-4 space-y-3">
      <h3 className="text-sm font-semibold">Obligations Timeline</h3>
      <div className="space-y-3 text-xs">
        {obligations.map((o) => (
          <div key={o.id} className="flex gap-2">
            <div className="mt-1 h-2 w-2 rounded-full bg-cyan-400" />
            <div>
              <p className="text-slate-100">{o.description}</p>
              {o.dueDate && (
                <p className="text-[10px] text-slate-400">
                  Due by {new Date(o.dueDate).toLocaleDateString()}
                </p>
              )}
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
