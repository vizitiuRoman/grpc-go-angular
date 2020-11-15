export function getUserIdFromJWT(jwt: string): number {
    const token = jwt.split('.')[1];
    const payload = JSON.parse(atob(token));

    return (payload as { userID: number }).userID;
}
