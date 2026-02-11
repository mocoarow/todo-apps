import { useEffect } from "react";
import { LoginPage } from "~/pages/LoginPage";
import { TopPage } from "~/pages/TopPage";
import { useAuthStore } from "~/stores/auth";

export default function App() {
  const { user, isLoading, fetchMe } = useAuthStore();

  useEffect(() => {
    fetchMe();
  }, [fetchMe]);

  if (isLoading) {
    return (
      <div className="flex min-h-svh items-center justify-center">
        <p className="text-muted-foreground">Loading...</p>
      </div>
    );
  }

  return user ? <TopPage /> : <LoginPage />;
}
