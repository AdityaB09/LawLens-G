import { getContracts } from "../../lib/api";
import ContractCard from "../../components/ContractCard";

export default async function ContractsPage() {
  const contracts = await getContracts();

  return (
    <div className="space-y-6">
      <header className="flex items-center justify-between">
        <h2 className="text-2xl font-semibold">Contracts</h2>
        <span className="text-sm text-slate-400">
          {contracts.length} saved contracts
        </span>
      </header>

      <div className="grid gap-4 md:grid-cols-2 xl:grid-cols-3">
        {contracts.map((c) => (
          <ContractCard key={c.id} contract={c} />
        ))}
      </div>
    </div>
  );
}
