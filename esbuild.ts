import fs from 'node:fs';

import * as esbuild from 'esbuild';
import {copy} from 'esbuild-plugin-copy';
import {sassPlugin} from 'esbuild-sass-plugin';
import postCssPlugin from 'esbuild-style-plugin';

const config: esbuild.BuildOptions = {
  entryPoints: ['resources/js/main.ts', 'resources/css/style.scss'],
  bundle: true,
  minify: false,
  logLevel: 'debug',
  metafile: true,
  sourcemap: true,
  legalComments: 'linked',
  allowOverwrite: true,
  outbase: 'resources',
  outdir: './assets/public/',
  target: ['es2020', 'chrome58', 'edge16', 'firefox57', 'safari11'],
  loader: {
    '.png': 'dataurl',
    '.woff': 'dataurl',
    '.woff2': 'dataurl',
    '.eot': 'dataurl',
    '.ttf': 'dataurl',
    '.svg': 'dataurl',
  },
  plugins: [
    sassPlugin({
      basedir: 'resources/css',
      loadPaths: ['node_modules', 'resources/css'],
    }),
    postCssPlugin({
      postcss: {
        plugins: [
          require('tailwindcss'),
          require('autoprefixer'),
        ],
      },
    }),
    copy({
      resolveFrom: 'cwd',
      assets: {
        from: ['./node_modules/htmx.org/dist/htmx.min.js'],
        to: ['./assets/public/js/vendor'],
      },
      watch: true,
    }),
    copy({
      resolveFrom: 'cwd',
      assets: {
        from: ['./node_modules/htmx.org/dist/ext/loading-states.js'],
        to: ['./assets/public/js/vendor'],
      },
      watch: true,
    }),
  ],
};

async function build() {
  try {  
    const result = await esbuild.build(config);
    fs.writeFileSync('meta.json', JSON.stringify(result.metafile));
  } catch (error) {
    console.error('esbuild error', error);
    process.exit(1);
  }
}

build();
