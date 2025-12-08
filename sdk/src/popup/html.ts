import type { ProviderResponse } from "../types.js";

export async function getOverviewHTML(providerData: ProviderResponse) {
    if (providerData.error !== null) {
        return `<p>${providerData.error}</p>`;
    }
    return `${providerData.data?.description ?? ""}`;
}

export async function getHTML(providerData: ProviderResponse) {
    const overview = await getOverviewHTML(providerData);
    const name = providerData.data?.display_name ?? "";
    const img = providerData.data?.image_url ?? "";
    return `
    <div class="integration-card" role="document">
      <div class="integration-card-header">
        <div class="left">
          <img src="${img}" alt="${name}" />
          <div>
            <h1 class="integration-card-title">${capitalize(name)}</h1>
            <p class="integration-card-subtitle">Where Work Happens</p>
          </div>
        </div>
        <div class="right">
          <button class="integration-connect-btn connective-connect-btn">Connect</button>
          <button class="connective-close-btn">✕</button>
        </div>
      </div>

      <div class="integration-tabs">
        <div class="integration-tab active connective-tab" data-tab="overview">Overview</div>
        <div class="integration-tab connective-tab" data-tab="configuration">Configuration</div>
      </div>

      <div class="integration-content connective-content-area">${overview}</div>
      <div class="integration-footer">Powered By Connective</div>
    </div>
  `;
}

export function capitalize(s: string) {
    return s ? s.charAt(0).toUpperCase() + s.slice(1) : "";
}
