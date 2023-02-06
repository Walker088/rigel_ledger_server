export type LoginPageInfo = {
    apiVersion: string,
    changeLogMd: string,
    ghOauthUrl?: string,

}

export type MainLayoutInfo = {
    locale: Locales,
};

export type JwtToken = {
    AccessToken: string,
    AccessTokenExpiry: number,
    RefreshToken: string,
    RefreshTokenExpiry: number,
};
