import type { LayoutServerLoad } from "./$types"
import type { MainLayoutInfo } from "$lib/types/home.type"

export const prerender = "auto";
export const ssr = false;

export const load: LayoutServerLoad<MainLayoutInfo> = async ({ locals: { locale } }) => {
    return { locale };
}
