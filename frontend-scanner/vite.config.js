import { build } from "vite";

// vite.config.js
export default {
    // config options
    build: {
      rollupOptions: {
        output: {
          // Define the entry point for the CSS
          entryFileNames: 'assets/[name].js',
          chunkFileNames: 'assets/[name].js',
          assetFileNames: ({ name }) => {
            // This condition ensures that CSS files are moved to the root directory
            if (name && name.endsWith('.css')) {
              return '[name][extname]';
            }
            return 'assets/[name][extname]';
          }
        }
      },
      // Directory where Vite will output the built files
      outDir: './../firmware-rpi/cmd/resources',
      // Ensure that the assets are placed in the same directory
      assetsDir: '.'
    }
  }