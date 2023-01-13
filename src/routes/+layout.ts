import type { LayoutLoad } from './$types'
import type { MainLayoutInfo } from "$lib/types/home.type"

import { loadLocaleAsync } from '$i18n/i18n-util.async'
import { setLocale } from '$i18n/i18n-svelte'

export const load: LayoutLoad<MainLayoutInfo> = async ({ data: { locale } }) => {
	await loadLocaleAsync(locale);

	setLocale(locale);

	return { locale }
}
