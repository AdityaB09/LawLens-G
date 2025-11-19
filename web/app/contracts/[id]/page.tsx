import { getContract, getContractClauses, getObligations } from "../../../lib/api";
import ClauseList from "../../../components/ClauseList";
import ClauseViewer from "../../../components/ClauseViewer";
import RiskHeatmap from "../../../components/RiskHeatmap";
import ObligationsTimeline from "../../../components/ObligationsTimeline";

type Props = {
  params: { id: string };
};

export default async function ContractDetailPage({ params }: Props) {
  const id = Number(params.id);
  const [contract, clauses, obligations] = await Promise.all([
    getContract(id),
    getContractClauses(id),
    getObligations(id)
  ]);

  return (
    <div className="grid gap-6 lg:grid-cols-[280px,minmax(0,1fr),320px]">
      <aside className="space-y-4">
        <h2 className="text-lg font-semibold">{contract.title}</h2>
        <p className="text-xs text-slate-400">
          {contract.partyA} ↔ {contract.partyB}
        </p>
        <ClauseList clauses={clauses} />
      </aside>

      <section className="space-y-4">
        <ClauseViewer clauses={clauses} />
        <RiskHeatmap clauses={clauses} />
      </section>

      <aside className="space-y-4">
        <ObligationsTimeline obligations={obligations} />
      </aside>
    </div>
  );
}
