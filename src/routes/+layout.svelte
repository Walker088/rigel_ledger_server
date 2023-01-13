<script lang="ts">
    import type { LayoutData } from './$types'
    import { setLocale } from '$i18n/i18n-svelte'

    import { page } from "$app/stores"
    import Fa from "svelte-fa"
    import { faGlobe } from "@fortawesome/free-solid-svg-icons"
    import "bootstrap/dist/js/bootstrap.bundle.min"
    import "../app.scss"

    export const prerender = true

    export let data: LayoutData;
    setLocale(data.locale);

    let selectedLang = data.locale;
    const replaceLocaleInUrl = (url: URL, locale: string, full = false): string => {
    	const [, , ...rest] = url.pathname.split('/')
    	const new_pathname = `/${[locale, ...rest].join('/')}`
    	if (!full) {
    		return `${new_pathname}${url.search}`
    	}
    	const newUrl = new URL(url.toString())
    	newUrl.pathname = new_pathname
    	return newUrl.toString()
    }
    const switchLang = () => {
        location.href = `${replaceLocaleInUrl($page.url, selectedLang, true)}`;
	};
</script>

<main>
    <nav class="navbar navbar-dark bg-dark">
        <div class="container-fluid">
            <div class="navbar-brand">
                <img src="/favicon.ico" width="30" height="30" alt="App Logo" class="align-text-top"> RigelLedger
            </div>
            <div class="nav-item d-inline-flex">
                <Fa color="dodgerblue" class="mt-2 me-2" icon={faGlobe}/>
                <select 
                    class="form-select form-select-sm" 
                    bind:value={selectedLang}
                    on:change={switchLang} 
                >
                    <option value="en">English</option>
                    <option value="zh">繁體中文</option>
                    <option value="es">Castellano</option>
                </select>
            </div>
        </div>
    </nav>

    <div class="container">
        <slot></slot>
    </div>
</main>
