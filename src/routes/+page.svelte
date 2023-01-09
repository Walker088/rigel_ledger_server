<script lang="ts">
	import Fa from "svelte-fa"
	import type { PageData } from './$types'
	import { faGoogle } from '@fortawesome/free-brands-svg-icons'

	export const prerender = true;
	export let data: PageData;
</script>

<svelte:head>
	<title>Home</title>
	<meta name="description" content="Rigel Ledger App home page"/>
	<meta name="author" content="walker088" />
</svelte:head>

<div class="grid-container">
	<section class="grid-item__login">
		<h1>Log in to your account</h1>
		<span><Fa icon={faGoogle} /> Sign In with Google</span>
	</section>
	<section class="grid-item__photo"></section>
	<section class="grid-item__info">
		<p class="h4">App Change Logs: <br>
			<span class="fs-6 fst-italic fw-light">API Version: {data.apiVersion}</span>
		</p>
		<p>{@html data.changeLogMd}</p>
	</section>
</div>

<style lang="scss">
	@use "$lib/styles/responsive";

	.grid-container {
		display: grid;
		grid-gap: 1rem;

		// Default Mobile Style
		grid-template-columns: 90vw;
		grid-template-rows: 20vh 45vh 30vh;
		.grid-item__login {
			grid-column: 1 / 2;
			grid-row: 1 / 2;
		}
		.grid-item__photo {
			grid-column: 1 / 2;
			grid-row: 2 / 3;
			background-image: url(/images/home/home.mobile.webp);
			background-size: contain;
			background-repeat: no-repeat;
			max-width: 90vw;
		}
		.grid-item__info {
			grid-column: 1 / 2;
			grid-row: 3 / 4;
		}

		@include responsive.base( responsive.$tablet-break ) {
			grid-template-columns: 90vw;
			grid-template-rows: 20vh 50vh 30vh;

			.grid-item__login {
				grid-column: 1 / 2;
				grid-row: 1 / 2;
			}

			.grid-item__photo {
				grid-column: 1 / 2;
				grid-row: 2 / 3;
				background-image: url(/images/home/home.desktop.webp);
				background-size: cover;
				background-repeat: no-repeat;
				max-width: 80vw;
			}

			.grid-item__info {
				grid-column: 1 / 2;
				grid-row: 3 / 4;
			}
		}

		@include responsive.base( responsive.$laptop-break ) {
			grid-template-columns: 40vw 60vw;
			grid-template-rows: repeat(4, 20vh);

			.grid-item__login {
				grid-column: 1 / 2;
				grid-row: 2 / 5;
			}

			.grid-item__photo {
				grid-column: 2 / 3;
				grid-row: 1 / 3;
				background-image: url(/images/home/home.desktop.webp);
				background-size: cover;
				background-repeat: no-repeat;
				max-width: 40vw;
			}

			.grid-item__info {
				grid-column: 2 / 3;
				grid-row: 3 / 5;
			}
		}
	}

</style>