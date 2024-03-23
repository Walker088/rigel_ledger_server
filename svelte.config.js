//import adapter from '@sveltejs/adapter-auto';
//import adapter from '@sveltejs/adapter-static';
import adapter from "svelte-adapter-bun";
//import { vitePreprocess } from '@sveltejs/kit/vite';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';
import * as child_process from 'node:child_process';

const svelteRoot = "svelte";

/** @type {import('@sveltejs/kit').Config} */
const config = {
	// Consult https://kit.svelte.dev/docs/integrations#preprocessors
	// for more information about preprocessors
	preprocess: vitePreprocess(),

	kit: {
		adapter: adapter(),
		alias: {
			$i18n: `src/i18n`,
			$lib: `src/${svelteRoot}/lib`,
			"$lib/*": `src/${svelteRoot}/lib/*`,
		},
		files: {
			assets: "static",
			hooks: {
				client: `src/${svelteRoot}/hooks.client`,
				server: `src/${svelteRoot}/hooks.server`,
			},
			lib: `src/${svelteRoot}/lib`,
			params: `src/${svelteRoot}/params`,
			routes: `src/${svelteRoot}/routes`,
			serviceWorker: `src/${svelteRoot}/service-worker`,
			appTemplate: `src/${svelteRoot}/app.html`,
			errorTemplate: `src/${svelteRoot}/error.html`,
		},
		version: {
			name: child_process.execSync('git rev-parse HEAD').toString().trim()
		}
	},
};

export default config;
