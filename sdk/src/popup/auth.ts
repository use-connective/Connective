export function openAuthWindow(opts: {
    baseURL: string;
    projectId: string;
    provider: string;
    userId: string;
    projectSecret: string;
}) {
    const { baseURL, projectId, provider, userId, projectSecret } = opts;
    const url = `${baseURL}/oauth/connect?provider=${encodeURIComponent(
        provider
    )}&projectID=${encodeURIComponent(projectId)}&userID=${encodeURIComponent(
        userId
    )}&projectSecret=${encodeURIComponent(projectSecret)}`;

    const popup = window.open(
        url,
        "authPopup",
        "width=500,height=600,left=100,top=100,resizable=no,scrollbars=yes"
    );

    if (!popup) {
        alert("Please enable popups");
    }

    return popup;
}
