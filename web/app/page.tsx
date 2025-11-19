import { getContracts } from "../lib/api";
import ContractCard from "../components/ContractCard";

export default async function DashboardPage() {
  const contracts = await getContracts();

  return (
    <div className="space-y-6">
      <header className="flex items-center justify-between">
        <div>
          <h2 className="text-2xl font-semibold">Contract Risk Overview</h2>
          <p className="text-sm text-slate-400">
            Upload contracts, map clauses, and explore risk like a SaaS product demo.
          </p>
        </div>
        <a
          href="/contracts"
          className="rounded-xl border border-cyan-500/40 bg-cyan-500/10 px-4 py-2 text-sm hover:bg-cyan-500/20"
        >
          View All Contracts
        </a>
      </header>

      <section className="grid gap-4 md:grid-cols-3">
        {contracts.slice(0, 3).map((c) => (
          <ContractCard key={c.id} contract={c} />
        ))}
      </section>
    </div>
  );
}
