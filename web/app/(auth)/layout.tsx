export default function AuthLayout({ children }: LayoutProps<"/">) {
  return (
    <div className="auth-page">
      <header className="arch">
        <div className="logo">foodies</div>
        <p className="tagline">Every shift, on record.</p>
      </header>
      <main className="panel auth-card">{children}</main>
    </div>
  );
}
