<script lang="ts">
    import { page } from '$app/stores';
	import LL from "$i18n/i18n-svelte";
</script>

<svelte:head>
	<title>{ $LL.home.pageTitle() }</title>
	<meta name="description" content="Rigel Ledger App error page"/>
	<meta name="author" content="walker088" />
</svelte:head>

<h1>{$page.status}: {$page?.error?.message || "Unknown Error Occured"}</h1>
