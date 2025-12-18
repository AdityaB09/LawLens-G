const API_BASE = process.env.NEXT_PUBLIC_API_BASE || "http://localhost:8080/api";

export type ContractSummary = {
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

export async function getContracts(): Promise<ContractSummary[]> {
  try {
    const res = await fetch(`${API_BASE}/contracts`, {
      cache: "no-store",
    });

    if (!res.ok) {
      console.error("API returned non-200:", res.status);
      return [];
    }

    const data = await res.json();
    return Array.isArray(data) ? data : [];
  } catch (err) {
    console.error("Failed to fetch contracts:", err);
    return [];
  }
}

export async function getContract(id: number) {
  const res = await fetch(`${API_BASE}/contracts/${id}`, { cache: "no-store" });
  if (!res.ok) throw new Error("Failed to fetch contract");
  return res.json();
}

export async function getContractClauses(id: number) {
  const res = await fetch(`${API_BASE}/contracts/${id}/clauses`, { cache: "no-store" });
  if (!res.ok) throw new Error("Failed to fetch clauses");
  return res.json();
}

export async function getObligations(id: number) {
  const res = await fetch(`${API_BASE}/contracts/${id}/obligations`, { cache: "no-store" });
  if (!res.ok) throw new Error("Failed to fetch obligations");
  return res.json();
}

export async function compareContracts(a: number, b: number) {
  const res = await fetch(`${API_BASE}/contracts/compare`, {
    method: "POST",
    headers: { "content-type": "application/json" },
    body: JSON.stringify({ contractA: a, contractB: b })
  });
  if (!res.ok) throw new Error("Failed to compare contracts");
  return res.json();
}
