<script lang="ts">
	import type { PageData } from "./$types"

	import LL from "$i18n/i18n-svelte"
	import Fa from "svelte-fa"
	import { faGithub } from "@fortawesome/free-brands-svg-icons"
	import { BreakPoints as bp } from "$lib/styles/responsive"

	export const prerender = true;
	export let data: PageData;

</script>

<svelte:head>
	<title>{ $LL.pageTitle() }</title>
	<meta name="description" content="Rigel Ledger App home page"/>
	<meta name="author" content="walker088" />
</svelte:head>

<div class="grid-container">
	<section class="grid-item__login">
		<h1>{ $LL.title() }</h1>
		<h4 class="fs-6 fw-light">{ $LL.subTitle() }</h4>
		<div class="d-inline-flex badge bg-secondary mt-4 fs-5">
			<Fa color="lightgrey" icon={faGithub} class="me-1"/>
			<span class="text-wrap"> Sign In with Github</span>
		</div>
	</section>
	<section class="grid-item__photo">
		<picture>
			<source media="(min-width:{bp.DesktopBreak}px)" srcset="/images/home/home.desktop.webp">
			<source media="(min-width:{bp.LaptopBreak}px)" srcset="/images/home/home.laptop.webp">
			<source media="(min-width:{bp.TabletBreak}px)" srcset="/images/home/home.tablet.webp">
			<img src="/images/home/home.mobile.webp" alt="Rigel in a jungle">
		</picture>
	</section>
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
		grid-template-columns: 85vw;
		grid-template-rows: 20vh 45vh 30vh;
		.grid-item__login {
			grid-column: 1 / 2;
			grid-row: 1 / 2;
		}
		.grid-item__photo {
			grid-column: 1 / 2;
			grid-row: 2 / 3;

			img {
				object-fit: contain;
				max-width: 100%;
    			max-height: 100%;
				width: 375px;
				height: 281px;
			}
		}
		.grid-item__info {
			grid-column: 1 / 2;
			grid-row: 3 / 4;
		}

		@include responsive.base( responsive.$tablet-break ) {
			grid-template-columns: 70vw;
			grid-template-rows: 20vh 50vh 30vh;
			.grid-item__login {
				grid-column: 1 / 2;
				grid-row: 1 / 2;
			}
			.grid-item__photo {
				grid-column: 1 / 2;
				grid-row: 2 / 3;
				img {
					object-fit: contain;
					max-width: 100%;
    				max-height: 100%;
					width: 768px;
					height: 576px;
				}
			}
			.grid-item__info {
				grid-column: 1 / 2;
				grid-row: 3 / 4;
			}
		}

		@include responsive.base( responsive.$laptop-break ) {
			grid-template-columns: 40vw 40vw;
			grid-template-rows: repeat(4, 20vh);
			.grid-item__login {
				grid-column: 1 / 2;
				grid-row: 2 / 5;
			}
			.grid-item__photo {
				grid-column: 2 / 3;
				grid-row: 1 / 4;
				img {
					object-fit: contain;
					max-width: 100%;
    				max-height: 100%;
					width: 969px;
					height: 726px;
				}
			}
			.grid-item__info {
				grid-column: 2 / 3;
				grid-row: 4 / 5;
			}
		}
	}

</style>