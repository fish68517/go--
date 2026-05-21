import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueDevTools from 'vite-plugin-vue-devtools'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    vueDevTools(),

  ],
    server: {
      port: 5178, // 确保端口和你的项目一致
      proxy: {
        // 配置代理规则
        '/api': {
          target: 'http://116.198.196.57:9000', // 后端接口地址
          changeOrigin: true, // 开启跨域
          rewrite: (path) => path.replace(/^\/api/, '') // 去掉路径中的/api前缀
        }
      }
    },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
  },
})
