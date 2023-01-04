<script lang="ts">
	import Fa from "svelte-fa"
	import type { PageData } from './$types'
	import { faGoogle } from '@fortawesome/free-brands-svg-icons'

	export const prerender = true;
	export let data: PageData;
	debugger;
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Rigel Ledger App home page"/>
	<meta name="author" content="walker088" />
</svelte:head>

<div class="container">
	<h1>Hello, Rigel!</h1>
	<Fa icon={faGoogle} />
	<p>{data.apiVersion}</p>
	<p>{data.changeLogMd}</p>
</div>

<style lang="scss">
	$font-stack: Helvetica, sans-serif;
	$primary-color: #333;

	.container {
		font: 100% $font-stack;
		color: $primary-color;
	}
</style>