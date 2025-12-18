import DiffView from "../../../components/DiffView";
export const dynamic = "force-dynamic";

export default function CompareContractsPage() {
  return (
    <div className="space-y-6">
      <header>
        <h2 className="text-2xl font-semibold">Compare Contract Versions</h2>
        <p className="text-sm text-slate-400">
          Enter two contract IDs to see how risk changes between versions.
        </p>
      </header>

      <DiffView />
    </div>
  );
}
