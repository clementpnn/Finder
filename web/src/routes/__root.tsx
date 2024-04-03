import NotFound from "@/components/errors/notFound";
import { createRootRoute, Outlet } from "@tanstack/react-router";

export const Route = createRootRoute({
  component: () => (
    <div className="h-screen flex mx-auto max-w-7xl p-4 sm:px-6 lg:px-8">
      <Outlet />
    </div>
  ),
  notFoundComponent: () => {
    return <NotFound />
  },
});
