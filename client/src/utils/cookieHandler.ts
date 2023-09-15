import Cookies from "js-cookie";

const userTokenCookieName = "userToken";

export function setUserToken(userToken: string): void {
  Cookies.set(userTokenCookieName, userToken);
}

export function getUserToken(): string | undefined {
  return Cookies.get(userTokenCookieName);
}
