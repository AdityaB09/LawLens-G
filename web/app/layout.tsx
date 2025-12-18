import "./globals.css";
import type { ReactNode } from "react";
export const dynamic = "force-dynamic";

export const metadata = {
  title: "LawLens-G",
  description: "AI Contract Clause Mapper & Risk Explorer"
};

export default function RootLayout({ children }: { children: ReactNode }) {
  return (
    <html lang="en">
      <body className="min-h-screen bg-slate-950 text-slate-100">
        <div className="flex min-h-screen">
          <aside className="w-64 border-r border-slate-800 bg-slate-900/80 p-4">
            <h1 className="text-xl font-semibold mb-4">LawLens-G</h1>
            <nav className="space-y-2 text-sm">
              <a href="/" className="block hover:text-cyan-400">Dashboard</a>
              <a href="/contracts" className="block hover:text-cyan-400">Contracts</a>
              <a href="/contracts/compare" className="block hover:text-cyan-400">Compare Versions</a>
            </nav>
          </aside>
          <main className="flex-1 p-6">{children}</main>
        </div>
      </body>
    </html>
  );
}
