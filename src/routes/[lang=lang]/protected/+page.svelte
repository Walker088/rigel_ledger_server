<script lang="ts">
    import type { PageData } from "./$types";
    import LL from "$i18n/i18n-svelte";

    export let data: PageData;
</script>

<svelte:head>
	<title>{ $LL.oauthRedirect.pageTitle() }</title>
	<meta name="description" content="Rigel Ledger App user main page"/>
	<meta name="author" content="walker088" />
</svelte:head>

<div class="grid-container">
    <h1>{$LL.protect.home.title( { userName: data.userName } )}</h1>
</div>

<style lang="scss">
    .grid-container {
        display: grid;
        grid-template-columns: 100vh;
        grid-template-rows: 100vw;
    }
</style>
