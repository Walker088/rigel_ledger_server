// See https://kit.svelte.dev/docs/types#app
// for information about these interfaces
// and what to do when importing types
/// <reference types="@sveltejs/kit" />

type Locales = import('$i18n/i18n-types').Locales;
type TranslationFunctions = import('$i18n/i18n-types').TranslationFunctions;

declare namespace App {
	// interface Error {}
	interface Locals {
        locale: Locales
		LL: TranslationFunctions
    }
	// interface PageData {}
	// interface Platform {}
}

