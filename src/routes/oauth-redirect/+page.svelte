<script lang="ts">
    import type { PageData } from "./$types";
    import { JwtManager } from '$lib/handlers/jwt';
    import LL from "$i18n/i18n-svelte";
	import { onMount } from "svelte";
    import Fa from "svelte-fa";
    //import { faGear } from "@fortawesome/free-solid-svg-icons";
    import { faGithub } from "@fortawesome/free-brands-svg-icons";

    export let jwtToken: PageData;

    onMount(() => {
        setTimeout(() => {
            JwtManager.setToken(jwtToken);
            window.location.href = `/${jwtToken.locale}/protected`;
        }, 3000);
    });
</script>

<svelte:head>
	<title>{ $LL.oauthRedirect.pageTitle() }</title>
	<meta name="description" content="Rigel Ledger App oauth page"/>
	<meta name="author" content="walker088" />
</svelte:head>

<div class="grid-container">
    <div class="grid-item">
        <Fa color="lightgrey" icon={faGithub} size="2x" flip="horizontal"/>
    </div>
</div>

<style lang="scss">
    .grid-container {
        display: grid;
        grid-template-columns: 100vh;
        grid-template-rows: 100vw;

        .grid-item {
            grid-column: 1 / 2;
            grid-row: 1 / 2;
            justify-self: stretch;
        }
    }
</style>
