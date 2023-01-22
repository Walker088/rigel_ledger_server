
export type LoginPageInfo = {
    apiVersion: string,
    changeLogMd: string,

    oauthOtps?: {
        github?: {
            title: string,
            link: string,
        },
        google?: {
            title: string,
            link: string,
        }
    }
}

export type MainLayoutInfo = {
    locale: Locales,
};
