export async function getHTML(providerData) {
    let overviewHTML = await getOverviewHTML(providerData);
    return `
    <div class="integration-card" role="document">

      <div class="integration-card-header">
        
        <!-- LEFT SECTION -->
        <div class="left">
          <img src="${providerData.data?.image_url}" alt="${providerData.data?.display_name}" />
          <div>
            <h1 class="integration-card-title">${capitalize(providerData.data?.display_name ?? '')}</h1>
            <p class="integration-card-subtitle">Where Work Happens</p>
          </div>
        </div>

        <!-- RIGHT SECTION (connect + close) -->
        <div class="right">
          <button class="integration-connect-btn connective-connect-btn">Connect</button>
          <button class="connective-close-btn">✕</button>
        </div>

      </div>

      <!-- TABS -->
      <div class="integration-tabs">
        <div class="integration-tab active connective-tab" data-tab="overview">Overview</div>
        <div class="integration-tab connective-tab" data-tab="configuration">Configuration</div>
      </div>

      <!-- CONTENT -->
      <div class="integration-content connective-content-area">
        ${overviewHTML}
      </div>

      <div class="integration-footer">Powered By Connective</div>
    </div>
  `;
}
export async function getOverviewHTML(providerData) {
    if (providerData.error !== null) {
        return `<p>${providerData.error} </p>`;
    }
    return `
    ${providerData.data?.description}
    `;
}
export function capitalize(s) {
    return s.charAt(0).toUpperCase() + s.slice(1);
}
//# sourceMappingURL=popup.html.js.map