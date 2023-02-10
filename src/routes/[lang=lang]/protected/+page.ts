import type { PageLoad } from "./$types";
import * as http from "$lib/handlers/api";
import { JwtManager } from "$lib/handlers/jwt";
import { error } from '@sveltejs/kit';

type UserInfo = {
    userId: string,
    userName: string,
    userMail: string,
    userTypeCode: number,
    userTypeName: string,
    mainCountry: string,
    mainCurrency: string,
    mainLanguage: Locales,

    ledgerSummary?: LedgerSummary,
}
type LedgerSummary = {
    balance: {
        assets: number,
        liabilityies: number,
        networth: number,
    },
    income: {
        monthlyIncome: number,
        monthlyExpence: number,
    },
    transactions: {
        availableRange: {
            firstAt: string,
            lastAt: string,
        },
        monthlyTxs: {
            transacId: number,
            transacDate: string,
            debitAccountCode: string,
            debitAccountName: string,
            creditAccountCode: string,
            creditAccountName: string,
            currency: string,
            originalAmt: number,
            adjustedAmt: number, // in user main currency
        }[],
    }
}

export const load: PageLoad<UserInfo> = async () => {
    const token = await JwtManager.getAccessToken() || undefined;
    if(!token) {
        throw error(401, `Unauthorized`);
    }
    const userId = JwtManager.getUidFromToken(token);
    if(!userId) {
        throw error(400, {message: ""});
    }
    const userInfo = await http.get(`v1/protected/user/${userId}`, [], token)
        .then(d => d as UserInfo);
    const ledgerSummary = await http.get(`v1/protected/ledger/${userId}/summary`, [], token)
        .then(d => d as LedgerSummary);
    userInfo.ledgerSummary = ledgerSummary;
    return userInfo;
};
