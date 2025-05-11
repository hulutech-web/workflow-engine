import { fileURLToPath, URL } from 'node:url'
import AutoImport from 'unplugin-auto-import/vite'
import { defineConfig,loadEnv } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import Components from 'unplugin-vue-components/vite';
import { AntDesignVueResolver } from 'unplugin-vue-components/resolvers';
export default defineConfig(({ command,mode }) => {
  const env = loadEnv(mode, process.cwd(), '')
  // console.log(env.VITE_API_URL)
  const isBuild = command == 'build'
  console.log(isBuild?'打包':'开发')
  return {
    esbuild: {
      drop: ['console', 'debugger'],
    },
    plugins: [
      vue(),
      vueJsx(),
      Components({
        resolvers: [
          AntDesignVueResolver({
            importStyle: false, // css in js
          }),
        ],
      }),
     
      AutoImport({
        //定义element-plus api按需加载
        resolvers: [AntDesignVueResolver()],
        imports: ['vue', 'vue-router'],
        //composables目录文件按需加载
        dirs: ['src/composables/**/*', 'src/enum/**/*', 'src/store/**/*', 'src/api/**/*'],
        dts: 'core/types/auto-imports.d.ts',
        vueTemplate: true,
        defaultExportByFilename: true,
      }),
    ],
    resolve: {
      alias: {
        '@': fileURLToPath(new URL('./src', import.meta.url))
      }
    },
    base: isBuild ? '/dist/' : '/',
    server: {
      host: true,
      proxy: {
        '/api': {
          target: `${env.VITE_API_URL}/api`,
          ws: true, //websocket支持
          changeOrigin: true,
          rewrite: (path) => path.replace(/^\/api/, '')
        },
      },
    },
  }
})
