export function getCSS(providerData) {
    return `
      /* Backdrop */
      .connective-backdrop {
        position: fixed;
        inset: 0;
        background: rgba(0,0,0,0.45);
        z-index: 2147483000;
        backdrop-filter: blur(1px);
      }

      /* Center popup */
      .connective-popup-container {
        position: fixed;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
        width: 100%;
        max-width: 600px;
        padding: 16px;
        z-index: 2147483001;
        box-sizing: border-box;
      }

      .integration-card {
        width: 100%;
        border-radius: 16px;
        background: #ffffff;
        border: 1px solid #ddd;
        box-shadow: 0 8px 18px rgba(0, 0, 0, 0.15);
        overflow: hidden;
        position: relative;
        font-family: system-ui, sans-serif;
      }

      /* Header */
      .integration-card-header {
        background: ${providerData.data?.primary_color};
        padding: 24px;
        display: flex;
        justify-content: space-between;
        align-items: center;
      }

      .integration-card-header .left {
        display: flex;
        gap: 16px;
        align-items: center;
      }

      .integration-card-header img {
        width: 56px;
        height: 56px;
        border-radius: 12px;
      }

      .integration-card-title {
        color: #fff;
        margin: 0;
        font-size: 24px;
        font-weight: 650;
      }

      .integration-card-subtitle {
        color: #ddd;
        margin-top: 4px;
        font-size: 14px;
      }

      /* RIGHT SIDE (desktop row: connect + close) */
      .integration-card-header .right {
        display: flex;
        align-items: center;
        gap: 8px;
      }

      /* Close button visible only on desktop */
      .connective-close-btn {
        background: transparent;
        border: none;
        font-size: 22px;
        cursor: pointer;
        color: white;
        opacity: 0.85;
      }

      /* Connect button */
      .integration-connect-btn {
        background: #fff;
        color: #3b0149;
        font-weight: 600;
        padding: 10px 22px;
        border-radius: 10px;
        border: none;
        cursor: pointer;
      }

      /* Tabs */
      .integration-tabs {
        display: flex;
        border-bottom: 1px solid #ddd;
        padding: 16px 24px 0 24px;
      }

      .integration-tab {
        padding: 10px 18px;
        border-radius: 8px 8px 0 0;
        cursor: pointer;
        font-size: 14px;
        margin-right: 12px;
        background: #eee;
        color: #444;
      }

      .integration-tab.active {
        background: #3b0149;
        color: #fff;
      }

      /* Content */
      .integration-content {
        padding: 24px;
        font-size: 15px;
        color: #444;
        max-height: 60vh;
        overflow-y: auto;
      }

      /* Footer */
      .integration-footer {
        text-align: center;
        font-size: 12px;
        padding: 14px;
        border-top: 1px solid #ddd;
        color: #aaa;
      }

      /* SCROLLBAR */
      .integration-content::-webkit-scrollbar {
        width: 6px;
      }
      .integration-content::-webkit-scrollbar-thumb {
        background: #ccc;
        border-radius: 6px;
      }

      @media (max-width: 640px) {

        /* Stack header vertically */
        .integration-card-header {
          flex-direction: column;
          align-items: flex-start;
          padding: 24px 20px;
          gap: 16px;
        }

        /* Make right side full width, stack connect button */
        .integration-card-header .right {
          width: 100%;
          display: block;
        }

        /* Connect button full-width */
        .integration-connect-btn {
          width: 100%;
          padding: 12px 18px;
          font-size: 18px;
          border-radius: 12px;
          text-align: center;
        }

        /* Hide close icon on mobile */
        .connective-close-btn {
          display: none;
        }

        /* Image smaller */
        .integration-card-header img {
          width: 48px;
          height: 48px;
        }

        .integration-card-title {
          font-size: 22px;
        }

        .integration-card-subtitle {
          font-size: 13px;
        }
      }`;
}
//# sourceMappingURL=popup.css.js.map