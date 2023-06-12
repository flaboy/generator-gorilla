import { defineConfig } from 'vite'
// import vue from '@vitejs/plugin-vue'
import path from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  root: path.resolve(__dirname, 'src'),
  resolve: {
    alias: {
      '~bootstrap': path.resolve(__dirname, 'node_modules/bootstrap'),
      '~bootstrap-icons': path.resolve(__dirname, 'node_modules/bootstrap-icons'),
    }
  },  
  // plugins: [vue()],
  build:{
    outDir: '../../public/',
    manifest: true,
    rollupOptions: {
      input: 'src/main.js',
    }
  }
})
