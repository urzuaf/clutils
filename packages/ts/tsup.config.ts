import { defineConfig } from 'tsup';

export default defineConfig({
  entry: ['src/index.ts', 'src/rut/index.ts', 'src/money/index.ts', 'src/date/index.ts', 'src/telecom/index.ts'],
  format: ['cjs', 'esm'],
  dts: true,
  clean: true,
  sourcemap: true,
  splitting: false,
  treeshake: true,
});
