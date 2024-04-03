import { defineConfig } from "vite"
import react from "@vitejs/plugin-react-swc"
import svgr from "vite-plugin-svgr"
import path from "node:path"
import { TanStackRouterVite } from "@tanstack/router-vite-plugin"
import tailwindcss from "@tailwindcss/vite"

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react(), svgr(), TanStackRouterVite(), tailwindcss()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src")
    }
  },
  preview: {
    port: 5173,
    strictPort: true,
    host: true
  }
})
